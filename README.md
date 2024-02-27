# Posts Web Application Backend

## Introduction

This project is the backend for a Posts Web Application, developed with Golang using the Gin framework and Gorm ORM for interacting with a PostgreSQL database. It provides a RESTful API for creating, reading, updating, and deleting (CRUD) post messages with titles.

## Features

- **CRUD Operations**: Supports creating, reading, updating, and deleting posts with titles.
- **RESTful API**: Offers a RESTful API for interaction with the frontend or any other client.
- **PostgreSQL Integration**: Uses PostgreSQL for storing and querying post data.
- **Scalable Architecture**: Designed with scalability in mind, making it easy to extend with more features.

## Prerequisites

Before you begin, ensure you have the following installed:

- Golang (version 1.15 or higher)
- PostgreSQL
- Git (for cloning the repository)

## Installation

Follow these steps to get the backend running:

```bash
# Clone the repository
git clone https://github.com/davidjohanhp/posts-app-backend.git

# Navigate to the project directory
cd posts-web-backend

# Install dependencies
go mod tidy

# Set up the PostgreSQL database
# Make sure PostgreSQL is running and create a database for this project

# Configure environment variables
# Create a .env file or export the following variables:
export DB_URL=yourdatabaseURL
```

## Running the Application

To start the server, run:

```bash
go run main.go
```

This will start the Gin server on the default port (usually `8080`), ready to handle API requests for managing posts.

## API Endpoints

The following endpoints are available:

- `POST /post` - Create a new post
- `GET /get-posts` - Retrieve all posts
- `GET /get-post/:id` - Retrieve a single post by its ID
- `PUT /update/:id` - Update a post by its ID
- `DELETE /post/:id` - Delete a post by its ID

## Building for Production

To build the application for production, run:

```bash
go build -o posts-web-backend
```

This command compiles the application into an executable named `posts-web-backend`.

## Contributing

Contributions are welcome! If you have suggestions for improving the application or have found bugs, please open an issue or submit a pull request.
