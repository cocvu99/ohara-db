# ohara-db

[![CI Pipeline](https://github.com/cocvu99/ohara-db/actions/workflows/ci.yaml/badge.svg)](https://github.com/cocvu99/ohara-db/actions)

## Table of Contents
- [About the Project](#about-the-project)
  - [Architecture Design](#architecture-design)
- [Getting Started](#getting-started)
- [Roadmap & Future Features](#roadmap-&-Future-Features)
- [Known Issues & Current Constraints](#Known-Issues-&-Current-Constraints)

## Project Overview
A relational database engine built from scratch using Go. This project is a hands-on implementation to deeply understand database internals, focusing on a disk-based B+ Tree storage engine, Key-Value interface, and concurrency control.


### Architecture Design

B+ Tree Data Structure Diagram
![B+ Tree Data Structure Diagram](docs/images/Bplus-tree.jpg)

<!-- *(Chỗ này sẽ cần 2-3 câu giải thích ngắn gọn về luồng hoạt động)* -->

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

## Roadmap & Future Features

This project is being developed in iterative milestones. Here is the current progress and future roadmap:

- [x] **Phase 1: In-Memory B+ Tree (Core Structure)**
  - [x] Define Internal Node and Leaf Node structures.
  - [x] Implement search and split logic.
  - [ ] Implement full insertion functionality.
- [ ] **Phase 2: Persistence (Disk-based B+ Tree)**
  - [ ] Implement Page/Disk Manager for 4KB blocks.
  - [ ] Build Free List Allocator for space reuse.
- [ ] **Phase 3: Key-Value Storage Engine**
  - [ ] Implement robust `Get`, `Set`, and `Del` operations.
- [ ] **Phase 4: Relational Model & Schema**
  - [ ] Map flat Key-Value to tables, rows, and columns.
  - [ ] Implement row serialization.
- [ ] **Phase 5: Advanced Querying & Indexing**
  - [ ] Implement iterators for Range Queries (`Seek`, `Next`).
  - [ ] Build Secondary Indexes for non-primary key lookups.
- [ ] **Phase 6: Concurrency & Transaction**
  - [ ] Implement ACID transactions.
  - [ ] Add Snapshot Isolation / Optimistic Concurrency Control.
- [ ] **Phase 7: SQL Query Execution**
  - [ ] Develop a SQL Parser and Expression Evaluator.


## Known Issues & Current Constraints

- **B+ Tree Search Logic Bug:** Currently, the `FindLastLE` function within the internal node search logic incorrectly identifies the *first* less-than-or-equal key instead of the *last* one. This is due to a premature `break` statement in the forward-iteration loop, which halts execution upon finding the first match. This impacts correct child-pointer routing when traversing the tree. The planned fix involves refactoring the loop to iterate backwards or removing the early break.
- **In-Memory Constraint:** The database currently resides entirely in RAM. Data will be lost upon application termination until the Disk-based Page Manager is fully implemented in Phase 2.