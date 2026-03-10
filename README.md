# School Management API

RESTful API for managing teachers, students, and subjects. Built with Go, Gin, and GORM.

## Architecture

Four-layer clean architecture:

- **Handler** - HTTP request/response handling
- **Service** - Business logic
- **Repository** - Database operations
- **Models** - Data structures

## Configuration

Copy `.env.sample` to `.env` and configure:

```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=school_management
SERVER_PORT=8080
```

## Running

```bash
go run cmd/api/main.go
```

## API Endpoints

All endpoints under `/api/v1`:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/teachers` | List all teachers |
| GET | `/teachers/:id` | Get teacher by ID |
| POST | `/teachers` | Create teacher |
| PUT | `/teachers/:id` | Update teacher |
| DELETE | `/teachers/:id` | Delete teacher |
| GET | `/students` | List all students |
| GET | `/students/:id` | Get student by ID |
| GET | `/students/:id/subjects` | Get student's enrolled subjects |
| POST | `/students` | Create student |
| PUT | `/students/:id` | Update student |
| DELETE | `/students/:id` | Delete student |
| GET | `/subjects` | List all subjects |
| GET | `/subjects/:id` | Get subject by ID |
| GET | `/subjects/:id/students` | Get students enrolled in subject |
| POST | `/subjects` | Create subject |
| PUT | `/subjects/:id` | Update subject |
| DELETE | `/subjects/:id` | Delete subject |
| GET | `/enrollments` | List all enrollments |
| GET | `/enrollments/:id` | Get enrollment by ID |
| POST | `/enrollments` | Enroll student in subject |
| PUT | `/enrollments/:id` | Update enrollment grade |
| DELETE | `/enrollments/:id` | Remove enrollment |

## Example Requests

**Create Teacher:**
```json
POST /api/v1/teachers
{
  "name": "Dr. Emily Brown",
  "email": "emily.brown@school.com",
  "subjectSpecialization": "Chemistry"
}
```

**Create Student:**
```json
POST /api/v1/students
{
  "name": "John Doe",
  "email": "john.doe@student.com",
  "enrollmentDate": "2026-01-15T00:00:00Z"
}
```

**Create Subject:**
```json
POST /api/v1/subjects
{
  "name": "Advanced Mathematics",
  "description": "Calculus and Linear Algebra",
  "teacherId": 1
}
```

**Enroll Student in Subject:**
```json
POST /api/v1/enrollments
{
  "studentId": 1,
  "subjectId": 1,
  "enrollmentDate": "2026-01-15T00:00:00Z",
  "grade": "A"
}
```

## Tech Stack

- Go 1.26.1
- Gin (HTTP framework)
- GORM (ORM)
- MySQL

