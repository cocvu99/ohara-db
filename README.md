# ohara-db

## Project Overview
A relational database engine built from scratch using Go. This project aims to implement a disk-based B+ Tree storage engine, a Key-Value interface, and essential database capabilities.

## Getting Started

### Prerequisites
- Go 1.24 or higher

### Build & Test
```bash
# Build the project
go build -v ./...

# Run unit tests
go test -v ./...
```

## Current Architecture
- **Language:** Go (1.24.5)
- **Module:** github.com/cocvu99/ohara-db
- **Core Structure:** In-memory B+ Tree

## Features Implemented
- **BTreeInternalNode**: Defines the structure for Internal Node with keys and children pointers.
- **Search Logic**: `FindLastLE` function to locate the correct insertion index.
- **Node Insertion**: `InsertKV` function to insert key-child pairs while maintaining array order.
- **Unit Testing**: Automated test coverage for B+ Tree insertion logic (`TestDatabase`).
- **CI/CD Integration**: GitHub Actions workflow (`ci.yaml`) configured for:
  - Automated `go build` and `go test` on `push` to `main` and `develop`.
  - Strict formatting (`gofmt`) and code analysis (`go vet`) on `pull_request` to `main`.

