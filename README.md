# URL Shortener API

A RESTful API service for creating, managing, and tracking shortened URLs. The application allows users to generate short links, manage their URLs, and securely access protected resources through JWT-based authentication.

---

# Features

- User Registration
- User Login
- JWT Authentication & Authorization
- Create Short URL
- Retrieve Original URL
- List User URLs
- Soft Delete URL
- Redis Caching
- PostgreSQL Database Integration


---

# Technology Stack

- [![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
- [![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?logo=go&logoColor=white)](https://gin-gonic.com/)
- [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
- [![Redis](https://img.shields.io/badge/Redis-FF4438?logo=redis&logoColor=white)](https://redis.io/)
- [![JWT](https://img.shields.io/badge/JWT-Auth-000000?logo=jsonwebtokens&logoColor=white)](https://jwt.io/)
- [![Swagger](https://img.shields.io/badge/Swagger-Docs-85EA2D?logo=swagger&logoColor=white)](https://swagger.io/)
- [![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

---


# Prerequisites

Before running the application, make sure you have installed:

- Go 1.24+
- PostgreSQL
- Redis
- Git

---

# Environment Variables

Create a `.env` file in the root directory.

```env
APP_PORT=<your port>

DB_HOST=<your db_hostt>
DB_PORT=<your db_port>
DB_USER=<your db_user>
DB_PASSWORD=<your db_pass>
DB_NAME=<your db_name>

JWT_SECRET=your-secret-key

REDIS_ADDR=<your rc_addr>
REDIS_PASS=<your rc_pass>
REDIS_USER=<your rc_user>

```

---

# Setup Instructions

## 1. Clone Repository

```bash
git clone https://github.com/iamhanif11/shortlink-backend.git

cd shortlink-backend
```

---

## 2. Install Dependencies

```bash
go mod tidy
```

---


---

## 3. Start PostgreSQL & Redis

- PostgreSQL
- Redis
- Migration
- Seeder

Jalankan:

```bash
docker compose up -d
```

Cek container:

```bash
docker ps
```


---

## 4. Run Database Migration with Makefile

Example:

```bash
make migrate-create NAME=<name_table>
```


---

## 5. Run Application

```bash
go run cmd/main.go
```

Server akan berjalan pada:

```text
http://localhost:8080
```

---

## 6.API Documentation

Swagger tersedia pada:

```text
http://localhost:8080/swagger/index.html
```

Generate ulang dokumentasi Swagger:

```bash
swag init -g cmd/main.go
```

---

## 📌 API Endpoints

### Authentication

| Method | Endpoint |
|----------|----------|
| POST | `/api/auth/register` |
| POST | `/api/auth/login` |
| DELETE | `/api/auth/logout` |

### Link Management

| Method | Endpoint |
|----------|----------|
| POST | `/api/links` |
| GET | `/api/links` |
| DELETE | `/api/links/:id` |

### Redirect

| Method | Endpoint |
|----------|----------|
| GET | `/:slug` |

---


# Assumptions & Design Decisions

### Authentication

- JWT is used for stateless authentication.
- Protected endpoints require a valid bearer token.

### Database

- PostgreSQL is used as the primary relational database.
- Soft delete is implemented using the `deleted_at` column.

### Caching

- Redis is used to cache URL lookup results.
- Cached data reduces database queries during redirects.

### URL Slug Generation

- Slugs are generated automatically.
- Slugs must be unique.
- Slug uniqueness is validated before insertion.

### Scalability

- Repository-Service-Handler architecture is used to separate concerns.
- Redis caching is introduced to improve read performance.
- Database migrations ensure consistent schema management.

---

# Future Improvements

- Custom slug support
- URL analytics
- Expiration date for links
- Rate limiting
- Refresh token mechanism
- Role-based authorization
- Docker Compose deployment

---

# Author

Developed as part of a Fullstack Web Development learning project.