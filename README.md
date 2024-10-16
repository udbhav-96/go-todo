# Go-Todo Application
This is a Todo Web Application built with Go. It demonstrates the use of clean architecture, routes, middleware, HTML templates, and database management.

## Project Structure

go-todo/ 
![Screenshot from 2024-10-16 21-45-44](https://github.com/user-attachments/assets/464e504c-192f-4a4c-b4d6-0dc1e67051f4)


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
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

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
