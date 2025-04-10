package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-github/github"
	pb "github.com/itachigit/github-search-service/githubsearchpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGithubSearchServiceServer
}

func (s *server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	query := req.SearchTerm
	if req.User != "" {
		query += fmt.Sprintf(" user:%s", req.User)
	}

	url := fmt.Sprintf("https://api.github.com/search/code?q=%s", url.QueryEscape(query))

	reqHttp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		reqHttp.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(reqHttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		var errorResp github.ErrorResponse

		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return nil, fmt.Errorf("error while decoding error response from github: %v", err.Error())
		}

		return nil, fmt.Errorf("error while requesting search term in github: %v, details: %v, refer: %v",
			errorResp.Message, errorResp.Errors[0].Message, errorResp.DocumentationURL)
	}

	var ghResp github.CodeSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&ghResp); err != nil {
		return nil, err
	}

	var results []*pb.Result
	for _, item := range ghResp.CodeResults {
		results = append(results, &pb.Result{
			FileUrl: *item.HTMLURL,
			Repo:    *item.Repository.FullName,
		})
	}

	return &pb.SearchResponse{Results: results}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGithubSearchServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
