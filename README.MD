# Book Store API

This is a simple Book Store API implemented in Go. It provides CRUD operations for managing books in a store using Gorilla Mux for routing.

## Features

- Add a new book
- Get a book by ID
- Update book details
- Delete a book
- List all books

## Project Structure

```sh
hello-gorld/
├── main.go
├── go.mod
├── models/
│   └── bookstore.go
├── handlers/
│   └── book_handlers.go
├── store/
│   └── bookstore.go //deprecated
│   └── books.json
├── utils/
│   └── update.go
└── middleware/
└── middleware.go
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- [Gorilla Mux](https://github.com/gorilla/mux)

### Installing

1. Clone the repository

```sh
git clone https://github.com/PaulinhoTortuga/hello-gorld.git
```

2. Install dependencies

```sh
 go mod tidy
```

3.  Running the Application

```sh
go run main.go
```

The server will start on http://localhost:8080.
