# Todo App Server

## Project Description
The Todo App Server is a Golang application built using the Gin framework, PostgreSQL for data storage, Docker for containerization, and Swagger for API documentation. This project provides CRUD (Create, Read, Update, Delete) operations for managing todo items, along with user authentication to secure the data.

### Features
- Create, Read, Update, and Delete todo items.
- User authentication for securing todo items.
- Documentation of API endpoints using Swagger.

## Table of Contents
1. [How to Install and Run the Project](#how-to-install-and-run-the-project)
2. [How to Use the Project](#how-to-use-the-project)

## How to Install and Run the Project
To run the Todo App Server locally, follow these steps:

### Prerequisites
Before you begin, ensure you have the following dependencies installed on your machine:
- Golang: [Installation Guide](https://golang.org/doc/install)
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- PostgreSQL: [Installation Guide](https://www.postgresql.org/download/)
- Git: [Installation Guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Clone the Repository
```bash
git clone https://github.com/beccoder/todo-app.git
cd todo-app
```

### Configure Environment Variables
Copy the `.env.example` file to `.env` and update database password `DB_PASSWORD` as needed.
```bash
cp .env.example .env
```

### Build and Run with Docker
Build and run the project using Docker Compose.
```bash
docker compose up
```
This command will start the Golang application and PostgreSQL database inside Docker containers and create needed database schema for project.

### Access the Application
The Todo App Server will be accessible at `http://localhost:8000`.

## How to Use the Project

### API Documentation

Swagger is used to document the API endpoints. You can access the API documentation at `http://localhost:8000/docs/index.html` once the application is running.

### User Authentication
To access certain API endpoints that require authentication, you will need to create a user account. Here is how:

1. Use an API client(e.g., Postman or cURL) to make a POST request to `http://localhost:8000/auth/sign-up` with a JSON body containing the following fields:
- `name`: Your desired name
- `username`: Your desired username
- `password`: Your desired password

2. After successfully registering, you can obtain an authentication token by making a POST request to `http://localhost:8000/auth/sign-in` with your username and password. This will return a JSON response containing an access token.
3. Include the access token in the `Authorization` header of your requests to authenticated endpoints. For example:
```bash
Authorization: Bearer your_access_token_here
```
### CRUD Operations
Use the provided API endpoints to perform CRUD operations on todo items. Refer to the Swagger documentation for details on each endpoint and their expected input and output.

## Badges

![GitHub contributors](https://img.shields.io/github/contributors/beccoder/todo-app.svg)
![GitHub stars](https://img.shields.io/github/stars/beccoder/todo-app.svg)
![GitHub forks](https://img.shields.io/github/forks/beccoder/todo-app.svg)
![GitHub issues](https://img.shields.io/github/issues/beccoder/todo-app.svg)
