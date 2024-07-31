# Go CMS Project

A simple CMS (Content Management System) built with Go, GORM, and the Gin framework.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Routes](#routes)
- [Docker](#docker)

## Features

- User Authentication (Login, Registration)
- JWT-based authentication middleware 
- Post creation, listing, and detail views 📝
- Commenting system
- Real-time updates with HTMX ⚡
- Responsive design with Bootstrap 📱
- TinyMCE advanced WYSIWYG HTML editor 🖋️
- Database management with Hasura using GraphQL with PostgreSQL 🗄️
- Containerization with Docker 🐳
- Post creation, listing, and detail views via admin
- Users creation, listing, and detail views via admin

## Technologies Used
- Backend: Go, PostgreSQL
- Frontend: HTMX, Templ, Flowbite, TinyMCE WYSIWYG HTML editor
- Containerization: Docker

## Requirements

- Go 1.22+
- A SQL database (e.g., PostgreSQL, MySQL, SQLite)
- Git

## Installation

1. **Clone the repository:**

```sh
   git clone https://github.com/Hrishikesh-Panigrahi/GoCMS.git
   cd GoCMS
   ```

2. **Install dependencies:**

```sh
   go mod tidy
   ```

And you are all set

## Usage

1. **Run the application:**

```sh
   air
   ```

2. **Open your browser and navigate to `http://localhost:8080`.**

## Docker

1. **Pull the Docker Image:**
```sh
docker pull hrishikeshpanigrahi025/my-go-app
```

2 . **Run the Docker container:**
```sh
docker run -p 8000:8000 hrishikeshpanigrahi025/my-go-app
```

3. **Open your browser and navigate to http://localhost:8000.**

## Project Structure
The project structure follows a standard Go project layout:

```
GoCMS/
├── connections/
├── controllers/
├── middleware/
├── models/
├── render/
├── routes/
├── services/
├── static/
├── templates/
├── Makefile
├── Dockerfile
├── main.go
└── go.mod
```

- connections/: Database connections
- controllers/: Handlers for the different routes
- middleware/: Handling middleware   
- models/: Database models
- render/: Render template files and redirection
- routes/: Route definitions
- services/: Helper functions
- static/: Static files (CSS, JS, images)
- templates/: templ templates
- Makefile: Makefile for automating tasks
- Dockerfile: Docker configuration
- main.go: Entry point of the application
- go.mod: Go module file

## Run Locally
To run the project locally, you have 3 options:

# Launch Debugger
1. Open your project in Visual Studio Code.
2. Set breakpoints as needed.
3. Launch the debugger by pressing F5 or by selecting Run > Start Debugging from the menu.

# Run Air
Ensure you have Air installed for live reloading.

1. Start Air by running the following command in your terminal:
```sh
air
```

# Run go run main.go Command
1. Open your terminal.

2. Navigate to the project directory.

3. Run the following command to start the application:
```sh
go run main.go
```

**Note:** Before running your project, make sure to generate the Templ files in the terminal to get the most updated UI. You can do this by running:
```sh
templ generate
```