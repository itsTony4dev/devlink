# DevLink

**DevLink** is a minimal and efficient backend service built in Go for managing developer bookmarks and code snippets. It provides a structured API for storing and retrieving curated development resources and snippets with tagging support.

## Features

- Clean and modular Go architecture
- Structured API endpoints for bookmarks and snippets
- Environment-based configuration
- Easily extendable project layout

## Project Structure

devlink/
├── cmd/
│ └── myapp/ # Application entry point
├── internal/
│ ├── config/ # Environment and configuration
│ ├── database/ # Database connection logic
│ ├── handlers/ # HTTP handler functions
│ ├── middleware/ # Custom middleware
│ └── models/ # Data models and logic
├── .env # Environment variables
├── .gitignore
├── go.mod
└── README.md


## Getting Started

### 1. Clone the Repository


git clone https://github.com/yourusername/devlink.git
cd devlink
2. Set Up Environment Variables
Create a .env file at the root:

env
PORT=8080
DB_URL=your_database_url
SECRET_KEY=your_secret_key
3. Run the Application
Using Air for live reloading:


go install github.com/air-verse/air@latest
air
Or run manually:

go run cmd/myapp/main.go