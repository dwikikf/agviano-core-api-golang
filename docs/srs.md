# 🧾 Software Requirements Specification (SRS)
## Proyek: Agviano Backend (Agviano Core API - RESTful API)
### Versi: 1.0
### Bahasa Pemrograman: Golang

---

## 🧭 1. Pendahuluan

### 1.1 Tujuan
Dokumen ini menjelaskan spesifikasi kebutuhan perangkat lunak untuk sistem backend **Agviano**, yang akan menjadi fondasi bagi:
- Website perusahaan (landing page),
- CMS internal untuk pengelolaan konten,
- Marketplace produk/jasa milik PT Agviano.

Backend ini akan menyediakan **RESTful API** yang digunakan oleh frontend berbasis **SvelteKit** dan potensi integrasi ke aplikasi mobile di masa depan.

### 1.2 Ruang Lingkup Sistem
Sistem backend Agviano bertugas untuk:
- Menyediakan API publik untuk menampilkan data di website (produk, artikel, halaman, kontak).
- Menyediakan API otentikasi dan otorisasi pengguna.
- Menyediakan API untuk mengelola konten (CMS).
- Menangani transaksi dan pemesanan produk (marketplace).
- Menyimpan dan mengelola data di basis data MySQL.

### 1.3 Pengguna Sistem
| Role | Deskripsi | Hak Akses |
|------|------------|-----------|
| **Admin** | Pengelola utama sistem (dari perusahaan) | CRUD semua data (produk, artikel, halaman, user, order, setting) |
| **Editor** | Pengelola konten CMS | CRUD artikel, halaman, media |
| **Customer** | Pengguna marketplace | Melihat produk, membuat pesanan, mengelola akun |
| **Public** | Pengunjung umum website | Mengakses data publik (produk, artikel, halaman) |

---

## 🧩 2. Kebutuhan Fungsional

### Modul A. Autentikasi & Manajemen User
- User dapat melakukan **registrasi akun**.
- User dapat melakukan **login** dan mendapatkan token JWT.
- Admin dapat melihat dan mengelola semua user.
- Sistem memvalidasi token pada setiap permintaan terautentikasi.
- Role user menentukan hak akses (middleware-based access control).

### Modul B. Manajemen Produk
- Admin dapat menambahkan, mengedit, dan menghapus produk.
- Setiap produk memiliki kategori.
- User (public) dapat melihat daftar dan detail produk.
- Mendukung upload gambar produk.
- Mendukung pagination dan filter (kategori, harga).

### Modul C. Manajemen Kategori
- Admin dapat membuat, mengubah, dan menghapus kategori.
- Produk dikaitkan dengan satu kategori.
- API publik untuk daftar kategori.

### Modul D. Transaksi / Order
- Customer dapat membuat order dari produk yang tersedia.
- Admin dapat mengubah status order (pending → paid → shipped).
- Sistem menyimpan data order dan item order.
- User dapat melihat riwayat order miliknya.
- Mendukung integrasi pembayaran (versi lanjutan).

### Modul E. Manajemen Artikel (Blog / News)
- Admin/Editor dapat membuat, mengedit, dan menghapus artikel.
- Artikel memiliki atribut: judul, slug, konten, cover image, tanggal publikasi.
- API publik untuk menampilkan daftar artikel dan detail artikel.

### Modul F. Manajemen Halaman Statis (Page)
- Admin/Editor dapat membuat halaman seperti *About*, *Contact*, *Privacy Policy*.
- API publik untuk mengambil halaman berdasarkan `slug`.

### Modul G. Manajemen Setting & Media
- Admin dapat menyimpan pengaturan global (logo, deskripsi perusahaan, kontak).
- Upload file (gambar, logo) disimpan di folder `uploads/` atau cloud storage.
- API untuk membaca konfigurasi publik (misal: logo website, alamat email, dsb).

---

## ⚙️ 3. Kebutuhan Non-Fungsional

| Aspek | Spesifikasi |
|--------|--------------|
| **Bahasa Pemrograman** | Go (Golang) |
| **Framework** | Gin (Web Framework) |
| **ORM** | GORM |
| **Database** | MySQL |
| **Auth** | JWT (JSON Web Token) |
| **API Type** | RESTful |
| **Performance** | Respon < 300ms per request di bawah 1000 RPS |
| **Availability** | 99% uptime |
| **Security** | Validasi input, hash password (bcrypt), JWT dengan secret key |
| **Scalability** | Dapat di-deploy terpisah dari frontend |
| **Documentation** | Swagger (auto-generated) |
| **Testing** | Unit test untuk usecase dan handler utama |

---

## 🧱 4. Arsitektur Sistem

### 4.1 Deskripsi Umum
Sistem akan menggunakan arsitektur **Clean Architecture** dengan 4 lapisan utama:
1. **Delivery (Interface Layer)** → Meng-handle HTTP request/respons.
2. **Use Case (Business Logic)** → Mengatur logika proses dan validasi.
3. **Repository (Data Access)** → Akses ke database (MySQL).
4. **Domain (Entities)** → Struktur data inti sistem.

### 4.2 Diagram Arsitektur

[SvelteKit Frontend]
↓
(HTTP)
↓
[ Gin HTTP Server ]
↓
[ UseCase Layer ]
↓
[ Repository Layer ]
↓
[ MySQL Database ]



---

## 🗃️ 5. Model Data (Konseptual)

### Entitas Utama

#### 🧑 User
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| name | string | Nama lengkap |
| email | string | Email unik |
| password_hash | string | Hash password |
| role | enum(admin, editor, customer) | Hak akses |
| created_at | datetime | Tanggal dibuat |

#### 🛍️ Product
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| name | string | Nama produk |
| description | text | Deskripsi produk |
| price | decimal | Harga produk |
| category_id | int | Relasi kategori |
| image_url | string | Gambar produk |
| created_at | datetime | Tanggal dibuat |

#### 🏷️ Category
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| name | string | Nama kategori |
| slug | string | Slug unik |
| created_at | datetime | Tanggal dibuat |

#### 📦 Order
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| user_id | int | User yang memesan |
| total_price | decimal | Total harga |
| status | enum(pending, paid, shipped) | Status order |
| created_at | datetime | Tanggal dibuat |

#### 📦 OrderItem
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| order_id | int | Relasi order |
| product_id | int | Relasi produk |
| quantity | int | Jumlah produk |
| price | decimal | Harga per item |

#### 📰 Article
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| title | string | Judul artikel |
| slug | string | Slug unik |
| content | text | Isi artikel |
| cover_image | string | Gambar utama |
| author_id | int | Relasi user penulis |
| published_at | datetime | Tanggal publikasi |

#### 📄 Page
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| title | string | Judul halaman |
| slug | string | Slug unik |
| content | text | Isi halaman |
| updated_at | datetime | Terakhir diubah |

#### ⚙️ Setting
| Field | Tipe | Deskripsi |
|--------|------|------------|
| id | int | Primary key |
| key | string | Nama pengaturan |
| value | text | Nilai pengaturan |

---

## 🔗 6. Integrasi & Antarmuka Eksternal

| Komponen | Fungsi |
|-----------|--------|
| **Frontend (SvelteKit)** | Mengonsumsi REST API |
| **File Storage / Cloud (opsional)** | Simpan media/gambar |
| **Payment Gateway (versi lanjutan)** | Proses pembayaran |
| **Swagger UI** | Dokumentasi API interaktif |

---

## 🧩 7. Risiko & Catatan Tambahan

| Risiko | Mitigasi |
|--------|-----------|
| Penggunaan JWT tanpa refresh token | Tambahkan mekanisme refresh token |
| Serangan SQL Injection / XSS | Gunakan parameterized query dan validator |
| Scaling server | Pisahkan layer API, DB, dan file storage |
| Akses tidak sah admin | Gunakan middleware otorisasi per role |

---

## 📅 8. Roadmap Implementasi

| Tahap | Modul | Target |
|--------|--------|--------|
| **1** | Setup project, config, DB connect | 1 minggu |
| **2** | Auth (register, login, JWT) | 1 minggu |
| **3** | Produk + Kategori | 1 minggu |
| **4** | Order | 1 minggu |
| **5** | Artikel + Page (CMS) | 1 minggu |
| **6** | Setting + Upload | 1 minggu |
| **7** | Swagger + Testing | 1 minggu |
| **8** | Deployment & Dokumentasi | 1 minggu |

---

## 📎 9. Referensi

- Golang Documentation: [https://go.dev/doc](https://go.dev/doc)
- Gin Framework: [https://gin-gonic.com/docs/](https://gin-gonic.com/docs/)
- GORM ORM: [https://gorm.io/docs/](https://gorm.io/docs/)
- Swagger Go: [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)

---

**Disusun oleh:**  
Dwiki Kautsar Fahmi  
**Tanggal:** 31 Oktober 2025  
**Versi:** 1.0
