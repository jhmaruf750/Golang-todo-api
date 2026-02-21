# 🚀 Golang Todo API

A RESTful Todo API built using **Golang**, **PostgreSQL**, **Chi Router**, and **JWT Authentication**.

This project demonstrates backend development concepts including clean project structure, authentication, middleware, and database integration.

---

## 📌 Features

- ✅ RESTful API (CRUD for Todos)
- ✅ PostgreSQL integration
- ✅ JWT Authentication
- ✅ Password Hashing (bcrypt)
- ✅ Protected Routes
- ✅ Chi Router
- ✅ Middleware (Logger & Recovery)
- ✅ Clean Project Structure

---

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Router:** Chi
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Password Hashing:** bcrypt

---

## 📁 Project Structure

```
todo-api/
│
├── main.go
│
├── database/
│   └── db.go
│
├── model/
│   ├── todo.go
│   └── user.go
│
├── handler/
│   ├── todo_handler.go
│   └── auth_handler.go
│
└── go.mod
```

---

## ⚙️ Installation & Setup

### 1️⃣ Clone Repository

```bash
git clone https://github.com/jhmaruf750/Golang-todo-api.git
cd Golang-todo-api
```

---

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 3️⃣ Setup PostgreSQL

Create database:

```sql
CREATE DATABASE todo_db;
```

Create tables:

```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
```

---

### 4️⃣ Update Database Connection

Update your PostgreSQL credentials inside:

```
database/db.go
```

---

### 5️⃣ Run Server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 🔐 Authentication Flow

### 📝 Register

```
POST /auth/register
```

Body:
```json
{
  "email": "test@mail.com",
  "password": "123456"
}
```

---

### 🔑 Login

```
POST /auth/login
```

Response:
```json
{
  "token": "your_jwt_token_here"
}
```

---

### 🔒 Access Protected Routes

Add header:

```
Authorization: Bearer YOUR_TOKEN
```

---

## 📚 API Endpoints

### Public Routes

| Method | Endpoint           | Description        |
|--------|-------------------|-------------------|
| GET    | /                 | Welcome message   |
| POST   | /auth/register    | Register user     |
| POST   | /auth/login       | Login user        |

---

### Protected Routes (JWT Required)

| Method | Endpoint          | Description        |
|--------|------------------|-------------------|
| GET    | /todos           | Get all todos     |
| POST   | /todos           | Create todo       |
| PUT    | /todos/{id}      | Update todo       |
| DELETE | /todos/{id}      | Delete todo       |

---

## 🧠 Learning Concepts Covered

- Go Modules
- Structs & JSON tags
- SQL queries in Go
- Router setup (Chi)
- Middleware usage
- JWT generation & validation
- Password hashing
- Clean code structure

---

## 🚀 Future Improvements

- Environment variable configuration (.env)
- Role-based authentication
- Refresh tokens
- Docker support
- Clean Architecture (Service & Repository layer)
- Validation layer
- Proper JSON error handling

---

## 👨‍💻 Author

**Jhmaruf750**

Backend Developer (Golang Enthusiast)

---

## 📜 License

MIT License
