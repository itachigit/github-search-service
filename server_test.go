package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/itachigit/github-search-service/githubsearchpb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var (
		srv        *server
		ctx        context.Context
		req        *githubsearchpb.SearchRequest
		mockServer *httptest.Server
	)

	BeforeEach(func() {
		srv = &server{}
		ctx = context.Background()

		mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("q") == "test user:example" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
                    "total_count": 1,
                    "items": [
                        {
                            "html_url": "https://github.com/example/repo/file",
                            "repository": {
                                "full_name": "example/repo"
                            }
                        }
                    ]
                }`))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"message": "Not Found"}`))
			}
		}))

		os.Setenv("GITHUB_API_URL", mockServer.URL)
	})

	AfterEach(func() {
		mockServer.Close()
	})

	Context("when a valid search request is made", func() {
		BeforeEach(func() {
			req = &githubsearchpb.SearchRequest{
				SearchTerm: "test",
				User:       "example",
			}
		})

		It("should return search results", func() {
			resp, err := srv.Search(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.Results).To(HaveLen(1))
			Expect(resp.Results[0].FileUrl).To(Equal("https://github.com/example/repo/file"))
			Expect(resp.Results[0].Repo).To(Equal("example/repo"))
		})
	})

	Context("when an invalid search request is made", func() {
		BeforeEach(func() {
			req = &githubsearchpb.SearchRequest{
				SearchTerm: "invalid",
				User:       "unknown",
			}
		})

		It("should return an error", func() {
			resp, err := srv.Search(ctx, req)

			Expect(err).ToNot(BeNil())
			Expect(resp).To(BeNil())
		})
	})

	Context("when the GitHub API returns an error", func() {
		BeforeEach(func() {
			req = &githubsearchpb.SearchRequest{
				SearchTerm: "error",
			}
		})

		It("should return an error", func() {
			resp, err := srv.Search(ctx, req)
			Expect(err).ToNot(BeNil())
			Expect(errors.Is(err, context.DeadlineExceeded)).To(BeFalse())
			Expect(resp).To(BeNil())
		})
	})
})
