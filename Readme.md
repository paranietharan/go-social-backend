# Go Social Backend

## Prerequisites
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (1.20 or higher recommended)
- [Air](https://github.com/cosmtrek/air) (for live reloading during development)

## Setup & Run

1. **Clone the repository** (if not already):
   ```bash
   git clone https://github.com/yourusername/go-social-backend.git
   cd go-social-backend
   ```

2. **Create a `.env` file** with your settings (example):
   ```bash
   echo "ADDR=:8080
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gosocial
DB_USER=postgres
DB_PASSWORD=postgres
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_MAX_IDLE_TIME=15m
DB_SSLMODE=disable
" > .env
   ```

3. **Build and run the database container**:
   ```bash
   docker compose up --build
   ```

4. **Run the server** with Air (for hot reloading):
   ```bash
   air
   ```
   Your server will be running at [http://localhost:8080](http://localhost:8080).

## Testing API Endpoints

Use the following curl commands to test the main endpoints:

### Health Check
```bash
curl http://localhost:8080/v1/health
```
Expected response:
```json
{"status":"ok"}
```

### Create User
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepassword",
    "first_name": "Test",
    "last_name": "User",
    "bio": "This is my bio",
    "avatar_url": "https://example.com/avatar.jpg"
  }'
```

### Create Post
```bash
curl -X POST http://localhost:8080/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "content": "This is the content of my first post",
    "user_id": 1,
    "tags": ["test", "golang", "api"]
  }'
```

These commands demonstrate basic functionality for creating users and posts. Adjust values as needed. For further customization, edit the `.env` file or modify the code in the [cmd/api](cmd/api) directory.````
<!-- filepath: /home/paranie/codes/go-social-backend/Readme.md -->

# Go Social Backend

## Prerequisites
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (1.20 or higher recommended)
- [Air](https://github.com/cosmtrek/air) (for live reloading during development)

## Setup & Run

1. **Clone the repository** (if not already):
   ```bash
   git clone https://github.com/yourusername/go-social-backend.git
   cd go-social-backend
   ```

2. **Create a `.env` file** with your settings (example):
   ```bash
   echo "ADDR=:8080
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gosocial
DB_USER=postgres
DB_PASSWORD=postgres
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_MAX_IDLE_TIME=15m
DB_SSLMODE=disable
" > .env
   ```

3. **Build and run the database container**:
   ```bash
   docker compose up --build
   ```

4. **Run the server** with Air (for hot reloading):
   ```bash
   air
   ```
   Your server will be running at [http://localhost:8080](http://localhost:8080).

## Testing API Endpoints

Use the following curl commands to test the main endpoints:

### Health Check
```bash
curl http://localhost:8080/v1/health
```
Expected response:
```json
{"status":"ok"}
```

### Create User
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepassword",
    "first_name": "Test",
    "last_name": "User",
    "bio": "This is my bio",
    "avatar_url": "https://example.com/avatar.jpg"
  }'
```

### Create Post
```bash
curl -X POST http://localhost:8080/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "content": "This is the content of my first post",
    "user_id": 1,
    "tags": ["test", "golang", "api"]
  }'
```

These commands demonstrate basic functionality for creating users and posts. Adjust values as needed. For further customization, edit the `.env` file or modify the code in the [cmd/api](cmd/api) directory.