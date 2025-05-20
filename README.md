# GitHub Manager API – Scalabit Challenge

This project implements a REST API in Go to manage GitHub repositories and pull requests, with a focus on clean code, proper structure, testing, and CI/CD.

## Features

- `POST /repos`: Create a new public repository
- `DELETE /repos/:owner/:repo`: Delete an existing repository
- `GET /repos`: List all repositories of the authenticated user
- `GET /repos/:owner/:repo/pulls?n=N`: List the N most recent open pull requests on a repository

## Tech Stack

- Language: **Go**
- Web Framework: **Gin**
- GitHub API Client: **go-github**
- Environment: **godotenv**
- Testing: **Go testing + Testify**
- CI/CD: **GitHub Actions** (to be added)

## Project Structure

```
github-manager/
├── cmd/                # Entry point (main.go)
├── internal/
│   ├── handlers/       # HTTP route handlers
│   └── services/       # GitHub API client logic
├── tests/              # Unit and integration tests
├── go.mod / go.sum     # Go dependencies
├── .env                # Local GitHub token (not committed)
└── README.md
```

## Authentication

You must provide a GitHub Personal Access Token in a `.env` file:

```env
GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxx
```

> You can generate one [here](https://github.com/settings/tokens) (Fine-grained access, with `repo` and `delete_repo` permissions).

## Running the API

```bash
go run cmd/main.go
```

The server runs by default at `http://localhost:8080`.

## Running Tests

```bash
go test ./tests -v
```

### Sample test coverage:
- Invalid input
- Missing route params
- Invalid query strings
- Token errors

## Example Requests

### Create a repo

```bash
curl -X POST http://localhost:8080/repos   -H "Content-Type: application/json"   -d '{"name": "my-new-repo"}'
```

### Delete a repo

```bash
curl -X DELETE http://localhost:8080/repos/<owner>/<repo>
```

### List pull requests

```bash
curl "http://localhost:8080/repos/<owner>/<repo>/pulls?n=3"
```

---


## 👤 Author

Created by [Mario (Mariola04)](https://github.com/Mariola04) as part of the Scalabit technical challenge.