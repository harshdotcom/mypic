# 📸 mypic — Backend API

Client → Router → Middleware → Controller → Service → Repository → Database → Response

A Go-based backend service built with Gin, GORM, and MySQL providing user authentication, profile management, and JWT-based authorization.

---

#  Tech Stack

- Go (1.21+ recommended)
- Gin (HTTP framework)
- GORM (ORM)
- MySQL
- JWT (authentication)
- bcrypt (password hashing)
- godotenv (env loading)
For Front end we are using angular
---

#  Mysql Queries Need to Run

CREATE DATABASE mypic_db;
USE mypic_db;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_name VARCHAR(100) NOT NULL UNIQUE,
  name VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT NULL UNIQUE,
  user_password LONGTEXT NOT NULL,
  user_logo_url LONGTEXT,
  time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE files (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  original_name VARCHAR(255) NOT NULL,
  stored_name VARCHAR(255) NOT NULL,
  extension VARCHAR(50),
  mime_type VARCHAR(100),
  size BIGINT,
  url TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---

#  Required Libraries

Install these before running:

```bash
go get github.com/gin-gonic/gin          # HTTP server
go get gorm.io/gorm                     # ORM
go get gorm.io/driver/mysql             # MySQL driver
go get golang.org/x/crypto/bcrypt       # Password hashing
go get github.com/golang-jwt/jwt/v5     # JWT auth
go get github.com/joho/godotenv          # Load env file





