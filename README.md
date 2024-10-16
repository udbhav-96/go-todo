# Go-Todo Application
This is a Todo Web Application built with Go. It demonstrates the use of clean architecture, routes, middleware, HTML templates, and database management.

## Project Structure

go-todo/
│
├── cmd/
│   └── web/
│       ├── handler.go        # Handlers for HTTP requests
│       ├── helpers.go        # Utility functions
│       ├── main.go           # Application entry point
│       ├── middleware.go     # Custom middleware for request handling
│       ├── routes.go         # API route definitions
│       └── templates.go      # Template parsing and rendering
│
├── internal/
│   └── models/
│       ├── errors.go         # Error handling logic
│       └── todo.go           # Todo model and database operations
│
├── ui/
│   ├── html/
│   │   ├── pages/
│   │   │   ├── home.tmpl     # Home page template
│   │   │   └── view.tmpl     # View todo items template
│   │   └── base.tmpl         # Base template layout
│   │
│   └── static/
│       ├── css/
│       │   └── main.css      # Styling for the web pages
│       └── img/              # Images directory
│
├── go.mod                    # Go module dependencies
├── go.sum                    # Dependency checksum file
└── projectStructure.txt      # Overview of project structure


## Prerequisites
Before setting up the database, ensure you have the following installed:

- Go (version 1.16 or higher)
- MySQL Server (version 5.7 or higher)
- A MySQL client (like MySQL Workbench, DBeaver, or command line)

## Database Setup

### Step 1: Create the Database

1. Open your MySQL client and log in to your MySQL server.
2. Run the following command to create the database:

   CREATE DATABASE gotodo;
   USE gotodo;

   CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);![Screenshot from 2024-10-16 21-45-44](https://github.com/user-attachments/assets/464e504c-192f-4a4c-b4d6-0dc1e67051f4)
![Screenshot from 2024-10-16 21-46-37](https://github.com/user-attachments/assets/c4168e48-3d49-4604-aec2-a09404d1cb1a)



## How to Run the Application / Installation

#### 1. Clone the repository:
git clone https://github.com/your-username/go-todo.git

#### 2. Navigate into the project directory:
cd go-todo

#### 3. Install dependencies:
go mod tidy

## Run the application:

Linux :
go run ./cmd/web

Windows: 
Google it.

### Open your browser and navigate to:
http://localhost:8080/todo

## Features
Add, view, and manage todo items.
Simple HTML templates for the UI.
Custom middleware for logging and error handling.
Database integration for persistent storage.

## Contributing
Feel free to open an issue or submit a pull request if you find any bugs or have suggestions for improvements.
