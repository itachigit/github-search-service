
# GitHub Search Service

A gRPC-based service for interacting with the GitHub Search API. This service allows users to search for code snippets in GitHub repositories using custom search terms and optional user filters.

## Features

- **Search GitHub Code**: Query the GitHub Search API for code snippets based on search terms.
- **Filter by User**: Narrow down search results to a specific GitHub user.
- **gRPC Interface**: Provides a robust and efficient gRPC-based API for clients.
- **Authentication**: Supports GitHub API authentication using personal access tokens.

---

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [API Definition](#api-definition)
- [Development](#development)
- [License](#license)

---

## Getting Started

This service is designed to be used as a backend for applications that need to search GitHub repositories programmatically. It uses gRPC for communication and requires a GitHub personal access token for authenticated requests.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/itachigit/github-search-service.git
   cd github-search-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Generate gRPC code (if needed):
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/github_search.proto
   ```

---

## Usage

### Running the Server

Start the gRPC server:
```bash
go run server.go
```

The server will start on port `50051` by default.

### Running the Client

Use the provided client to test the service:
```bash
go run client.go
```

---

## Environment Variables

The service requires the following environment variable:

- **`GITHUB_TOKEN`**: A GitHub personal access token for authenticated requests. This is optional but recommended to avoid rate limits.

Set the environment variable before running the server:
```bash
export GITHUB_TOKEN=your_personal_access_token
```

---

## API Definition

### gRPC Service: `GithubSearchService`

#### RPC Methods

1. **Search**
   - **Request**: `SearchRequest`
     - `search_term` (string): The term to search for.
     - `user` (string): (Optional) GitHub username to filter results.
   - **Response**: `SearchResponse`
     - `results` (array): List of search results.
       - `file_url` (string): URL to the file in the repository.
       - `repo` (string): Full name of the repository.

#### Example Request
```proto
message SearchRequest {
  string search_term = 1;
  string user = 2;
}
```

#### Example Response
```proto
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string file_url = 1;
  string repo = 2;
}
```

---

## Development

### Prerequisites

- Go 1.20 or later
- Protocol Buffers Compiler (`protoc`) with Go plugins:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

### Generating Protobuf Code

If you make changes to the `.proto` file, regenerate the gRPC code:
```bash
protoc --go_out=. --go-grpc_out=. github_search.proto
```

### Running Tests

Add unit tests for your code and run them using:
```bash
go test ./...
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.

---

## Acknowledgments

- [GitHub API](https://docs.github.com/en/rest)
- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
```

### Key Improvements:
1. **Detailed Features**: Clearly outlines what the service does.
2. **Installation and Usage**: Step-by-step instructions for setup and running the service.
3. **Environment Variables**: Explains the required GitHub token.
4. **API Definition**: Provides a clear description of the gRPC service and its methods.
5. **Development Section**: Includes instructions for generating protobuf code and running tests.
6. **Contributing**: Encourages collaboration with clear steps.
7. **Acknowledgments**: Credits the tools and APIs used in the project.

This README.md should provide a comprehensive overview of your project and make it easy for others to understand and contribute.