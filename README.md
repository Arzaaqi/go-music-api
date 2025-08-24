# 🎶 Go Music API - Belajar Framework GIN
Halo teman-teman semuanya, makin penasaran sama ngoding dengan golang. Kali ini saya membuat RESTful API sederhana untuk mengelola koleksi musik, dibangun menggunakan **Golang**, **Gin**, dan **GORM**.  
Proyek ini dibuat untuk latihan memahami konsep **Backend Golang Framework** dan integrasi dengan **database MySQL**.

---

## ✨ Features
- CRUD (Create, Read, Update, Delete) data musik
- Swagger API Documentation
- Database integration dengan MySQL menggunakan GORM
- Clean project structure (Controller → Service → Repository → Model → Config)
- Auto migration untuk model

---

## 📂 Project Structure
go-music-api/

│── config/ # konfigurasi database

│── controllers/ # HTTP request handler

│── docs/ # swagger docs (auto generated)

│── models/ # model (entity) GORM

│── repositories/ # database access logic

│── services/ # business logic

│── routes/ # route handler

│── main.go # entry point

---

## 🚀 Installation

### 1. Clone repository
```bash
git clone https://github.com/username/go-music-api.git
cd go-music-api
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Copy .env.example menjadi .env dan isi sesuai konfigurasi MySQL:
```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=gomusic
DB_USERNAME=root
DB_PASSWORD=
```

### 4, Jalankan aplikasi
```bash
swag init

go run main.go
```

---

## 📖 API Documentation (Swagger)

Setelah aplikasi berjalan, buka:
```bash
http://localhost:8080/swagger/index.html
```

---

## 📌 API Endpoints
🎵 Musics
```bash

GET /musics → Ambil semua musik

GET /musics/{id} → Ambil musik berdasarkan ID

POST /musics → Tambah musik baru

PUT /musics/{id} → Update musik

DELETE /musics/{id} → Hapus musik

```

---

## 🛠️ Tech Stack
- Go
- Gin
- GORM
- Swagger

---

# SEMANGAT BELAJAR 
