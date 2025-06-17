# DevLink 🚀

DevLink is a powerful platform for developers to save, organize, and share their favorite coding resources. Built with Go, it provides a robust API for managing bookmarks and code snippets with advanced features like tagging, search, and organization.

## Features ✨

- **Resource Management**
  - Save and organize coding resources (articles, GitHub repos, tools)
  - Add descriptions and tags to resources
  - Edit and delete resources
  - Search resources by title, description, or URL
  - Filter resources by tags

- **User Management**
  - Secure user authentication with JWT
  - User registration and login
  - Profile management
  - Resource ownership and privacy

- **Advanced Features**
  - Pagination for large datasets
  - Full-text search
  - Tag-based filtering
  - Rate limiting
  - CORS support
  - Security headers

## Tech Stack 🛠

- **Backend**: Go (Golang)
- **Database**: SQLite (with GORM)
- **Authentication**: JWT
- **API**: RESTful
- **ORM**: GORM
- **Router**: Gorilla Mux
- **Configuration**: godotenv

## Project Structure 📁

```
devlink/
├── cmd/                    # Application entry points
│   └── devlink/           # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── db/              # Database connection
│   ├── dto/             # Data Transfer Objects
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── repository/      # Data access layer
│   ├── routes/          # Route definitions
│   └── utils/           # Utility functions
├── .env                  # Environment variables
├── .gitignore
├── go.mod
└── README.md
```

## Getting Started 🚀

### Prerequisites

- Go 1.24 or higher
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/itsTony4dev/devlink.git
   cd devlink
   ```

2. Set up environment variables:
   Create a `.env` file in the root directory:
   ```env
   PORT=8080
   JWT_SECRET=your_jwt_secret
   DB_URL=devlink.db
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run cmd/devlink/main.go
   ```

## API Endpoints 📡

### Authentication
```
POST /users/register    - Register a new user
POST /users/login      - Login user
POST /users/logout     - Logout user
```

### User Management
```
GET    /users          - Get all users (paginated)
GET    /users/{id}     - Get user by ID
PUT    /users/{id}     - Update user
DELETE /users/{id}     - Delete user
```

### Resources
```
POST   /resources           - Create a new resource
GET    /resources           - Get user's resources (paginated)
GET    /resources/{id}      - Get a specific resource
PUT    /resources/{id}      - Update a resource
DELETE /resources/{id}      - Delete a resource
GET    /resources/search    - Search resources
GET    /resources/tags      - Get resources by tags
```

## API Examples 📝

### Create a Resource
```bash
curl -X POST http://localhost:8080/resources \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go Best Practices",
    "url": "https://example.com/go-best-practices",
    "description": "A comprehensive guide to Go best practices",
    "tags": ["go", "best-practices", "programming"]
  }'
```

### Search Resources
```bash
curl -X GET "http://localhost:8080/resources/search?q=go&page=1&pageSize=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Security 🔒

- JWT-based authentication
- Password hashing with bcrypt
- Rate limiting
- CORS configuration
- Security headers
- Input validation
- Resource ownership validation

## Best Practices Implemented ✅

- Clean Architecture
- Repository Pattern
- DTO Pattern
- Middleware Pattern
- Error Handling
- Dependency Injection
- Configuration Management
- Database Migrations
- API Documentation
- Code Organization


## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙏

- [Gorilla Mux](https://github.com/gorilla/mux)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)


