# MyPOS Backend API (Golang + Gin Framework)

Service backend REST API lokal berbasis **Gin Web Framework** untuk aplikasi POS Toko Kelontong standalone (100% offline-first).

---

## Struktur Folder

```
backend/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point aplikasi backend
├── internal/
│   ├── router/
│   │   └── router.go               # Pendaftaran endpoint/URL routing API (Gin)
│   ├── config/
│   │   └── config.go              # Pemuatan environment & konfigurasi
│   ├── database/
│   │   ├── db.go                  # Koneksi pool MySQL
│   │   └── migrations/
│   │       └── 000001_init_schema.sql # Skema tabel database (Users, Products, Transactions, dll)
│   ├── handler/                   # HTTP Handlers / Controllers
│   ├── middleware/
│   │   └── cors.go                # Middleware CORS untuk komunikasi frontend Next.js lokal
│   ├── model/                     # Struct Entitas domain & DTO (User, Product, Transaction, dll)
│   ├── repository/                # Layer query SQL database
|   |── router/                    # Routing API
│   └── service/                   # Layer logika bisnis (ACID Checkout, Stok Alert, dll)
├── pkg/
│   └── response/
│       └── response.go            # Helper standarisasi respon JSON API
├── .env.example                   # Contoh konfigurasi environment
├── go.mod                         # Definisi modul Go & dependensi
└── README.md
```

---

## Cara Menjalankan

1. **Pastikan MySQL lokal aktif**:
   Buat database bernama `mypos_db`:
   ```bash
   mysql -u root -e "CREATE DATABASE IF NOT EXISTS mypos_db;"
   ```
   Jalankan migrasi skema tabel:
   ```bash
   mysql -u root mypos_db < internal/database/migrations/000001_init_schema.sql
   ```

2. **Salin konfigurasi environment**:
   ```bash
   cp .env.example .env
   ```

3. **Jalankan Server Go**:
   ```bash
   go run cmd/api/main.go
   ```
   Server akan berjalan di `http://localhost:8080`.
   Endpoint cek kesehatan: `GET http://localhost:8080/api/health`.
