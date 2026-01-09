Client → Router → Middleware → Controller → Service → Repository → Database
                                     ↓
                                  Response ←––––––––––––––––––––––––––––––

mypic/
│── cmd/                # Application entry point
│── config/             # DB and JWT config
│── controllers/        # HTTP handlers
│── services/           # Business logic
│── repositories/       # Database operations
│── models/             # DB models
│── routes/             # API route definitions
│── middlewares/        # Auth middleware
│── .env                # Environment variables
│── go.mod / go.sum


# 📸 mypic — Backend API

A Go-based backend service built with **Gin**, **GORM**, and **MySQL** providing user authentication, profile management, and JWT-based authorization.

---

## 🚀 Tech Stack

- Go (1.21+ recommended)
- Gin (HTTP framework)
- GORM (ORM)
- MySQL
- JWT (authentication)
- bcrypt (password hashing)

---

## 📁 Project Structure

