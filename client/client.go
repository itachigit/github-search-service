package main

import (
	"context"
	"log"
	"time"

	pb "github.com/itachigit/github-search-service/githubsearchpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGithubSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := &pb.SearchRequest{
		SearchTerm: "outputsize",
		User:       "torvalds",
	}

	resp, err := client.Search(ctx, req)
	if err != nil {
		log.Fatalf("Search RPC failed: %v", err)
	}

	for _, result := range resp.Results {
		log.Printf("Found file: %s, in repo: %s", result.FileUrl, result.Repo)
	}
}
