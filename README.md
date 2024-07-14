# Go CMS Project

A simple CMS (Content Management System) built with Go, GORM, and the Gin framework.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Routes](#routes)
- [Contributing](#contributing)

## Features

- User Authentication (Login, Registration)
- CRUD operations for posts
- Light/Dark mode support for the login page
- Responsive design using Tailwind CSS

## Requirements

- Go 1.22+
- A SQL database (e.g., PostgreSQL, MySQL, SQLite)
- Git

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/Hrishikesh-Panigrahi/GoCMS.git
   cd GoCMS
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up your database and update the database connection details in the `config` file.

4. Run database migrations:

   ```sh
   go run main.go migrate
   ```

## Usage

1. Run the application:

   ```sh
   go run main.go
   ```

2. Open your browser and navigate to `http://localhost:8080`.

## Project Structure

```
gocms/
│
├── controllers/
│ ├── post_controller.go
│ └── user_controller.go
│
├── models/
│ ├── post.go
│ └── user.go
│
├── templates/
│ ├── index.html
│ └── login.html
│
├── static/
│ ├── css/
│ │ └── styles.css
│ └── js/
│ └── scripts.js
│
├── main.go
├── go.mod
└── go.sum
```

## Routes

- `GET /` - Home page, list all posts
- `GET /post/:id` - View a single post
- `GET /login` - Login page
- `POST /login` - Login action
- `GET /register` - Registration page
- `POST /register` - Registration action
- `GET /post/new` - New post page
- `POST /post` - Create a new post
- `GET /post/edit/:id` - Edit post page
- `POST /post/update/:id` - Update post action
- `POST /post/delete/:id` - Delete post action

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.
