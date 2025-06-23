# Students API

A simple RESTful API built with Go for managing student records.

## Features

- **Add new students:** Create a new student record.
- **Retrieve a single student by ID:** Fetch information about a specific student using their unique ID.
- **Retrieve a list of all students:** Get a list of all student records.

## Technologies Used

- **Language:** Go (Golang)
- **Framework:** net/http (standard library)
- **Database:** SQLite

## Getting Started

### Prerequisites

- Go 1.18 or later installed
- SQLite installed (or use the default SQLite file database)

### Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/students-api.git
    cd students-api
    ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```

3. **Set up the database:**
    - The API will automatically create a SQLite database file (e.g., `students.db`) if it doesn't exist.
    - If you need to run migrations manually, provide those instructions here.

4. **Run the server:**
    ```bash
    go run main.go
    ```
    Or, to build and run:
    ```bash
    go build -o students-api
    ./students-api
    ```

By default, the API will run on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint           | Description                       |
|--------|--------------------|-----------------------------------|
| POST   | `/students`        | Add a new student                 |
| GET    | `/students`        | Get list of all students          |
| GET    | `/students/{id}`   | Get a single student by ID        |

### Example Requests

**Add a new student**
```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 20, "email": "john@example.com"}'
```

**Get all students**
```bash
curl http://localhost:8080/students
```

**Get a student by ID**
```bash
curl http://localhost:8080/students/1
```

## Authentication

No authentication is required for this API.

## License

This project is licensed under the MIT License.

---

Feel free to contribute or raise issues!
