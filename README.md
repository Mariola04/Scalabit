
# Scalabit GitHub API Challenge

This project implements a RESTful API in **Go** that interacts with the GitHub API to:

- Create repositories
- Delete repositories
- List repositories
- List N open pull requests on a specific repository

It also includes a full CI/CD pipeline using **GitHub Actions**, with:

- Unit tests
- Linting (`staticcheck`)
- Security scanning (`gosec`)
- Deployment placeholder

---

## How to run locally

1. **Set your GitHub token**

Create a `.env` file:

```env
GITHUB_TOKEN=your_personal_access_token
```

> Make sure the token has repo permissions.

2. **Run the project**

```bash
go run cmd/main.go
```

3. **Use curl or Postman** to call:

- `GET /repos` → list your repositories
- `POST /repos` with JSON `{ "name": "my-repo" }` → create a repository
- `DELETE /repos/:owner/:repo` → delete a repository
- `GET /repos/:owner/:repo/pulls?n=3` → list 3 open PRs

---

## GitHub Actions CI/CD

This repo includes a complete pipeline under `.github/workflows/go.yml`.

It performs:

- `go test`
- `staticcheck`
- `gosec`
- Deployment step on main branch (currently a placeholder)

---

## Tests

Tests are located under `/tests` and can be run locally with:

```bash
go test ./tests -v
```

---

## Security

I used [gosec](https://github.com/securego/gosec) to catch security issues like:

- Unhandled errors
- Insecure function use
- Credential leakage

---

## Status

![Go](https://github.com/Mariola04/Scalabit/actions/workflows/go.yml/badge.svg)

---

## Structure

```
.
├── cmd/                 # main entry point
├── internal/
│   ├── handlers/        # HTTP handlers
│   └── services/        # GitHub API logic
├── tests/               # Test files
├── .github/workflows/   # CI pipeline
└── README.md
```

---


## Author

Created by Mario (Mariola04) as part of the Scalabit technical challenge.