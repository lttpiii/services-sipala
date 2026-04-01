# 📚 Database Schema – ToolTrack

## 🔹 Extensions

```sql
uuid-ossp → untuk generate UUID otomatis
```

---

## 🔹 Enums

### user_role

```text
admin → full akses
staff → approve & monitoring
borrower → peminjam
```

### borrow_status

```text
pending → menunggu approval
approved → disetujui
rejected → ditolak
returned → sudah dikembalikan
```

---

# 👤 users

| Field         | Type        | Constraint       | Keterangan           |
| ------------- | ----------- | ---------------- | -------------------- |
| id            | uuid        | PK               | ID unik user         |
| name          | varchar     | NOT NULL         | Nama user            |
| email         | varchar     | UNIQUE, NOT NULL | Untuk login          |
| password_hash | text        | NOT NULL         | Password terenkripsi |
| role          | enum        | NOT NULL         | Role user            |
| created_at    | timestamptz | NOT NULL         | Waktu dibuat         |
| updated_at    | timestamptz | NULL             | Waktu update         |
| deleted_at    | timestamptz | NULL             | Soft delete          |

💡 **Catatan:**

- `deleted_at` untuk menjaga data tetap ada (audit)
- `email` wajib unique untuk login system

---

# 🏷️ categories

| Field      | Type        | Constraint | Keterangan    |
| ---------- | ----------- | ---------- | ------------- |
| id         | uuid        | PK         | ID kategori   |
| name       | varchar     | NOT NULL   | Nama kategori |
| created_at | timestamptz | NOT NULL   | Waktu dibuat  |
| updated_at | timestamptz | NULL       | Waktu update  |

---

# 🧰 tools

| Field       | Type        | Constraint | Keterangan         |
| ----------- | ----------- | ---------- | ------------------ |
| id          | uuid        | PK         | ID alat            |
| name        | varchar     | NOT NULL   | Nama alat          |
| category_id | uuid        | FK         | Relasi ke kategori |
| stock       | int         | CHECK ≥ 0  | Jumlah stok        |
| description | text        | NULL       | Deskripsi          |
| created_at  | timestamptz | NOT NULL   | Waktu dibuat       |
| updated_at  | timestamptz | NULL       | Waktu update       |
| deleted_at  | timestamptz | NULL       | Soft delete        |

💡 **Catatan:**

- `stock` tidak boleh negatif
- relasi ke kategori menjaga konsistensi data

---

# 📥 borrow_transactions (Header Transaksi)

| Field       | Type        | Constraint          | Keterangan           |
| ----------- | ----------- | ------------------- | -------------------- |
| id          | uuid        | PK                  | ID transaksi         |
| borrower_id | uuid        | FK                  | User peminjam        |
| status      | enum        | DEFAULT pending     | Status transaksi     |
| borrow_date | timestamptz | NOT NULL            | Tanggal mulai pinjam |
| due_date    | timestamptz | CHECK ≥ borrow_date | Batas pengembalian   |
| approved_by | uuid        | FK, NULL            | Petugas yang approve |
| approved_at | timestamptz | NULL                | Waktu approval       |
| created_at  | timestamptz | NOT NULL            | Waktu dibuat         |
| updated_at  | timestamptz | NULL                | Waktu update         |

💡 **Catatan penting:**

- ini adalah **header transaksi**
- 1 transaksi bisa punya banyak item

---

# 📦 borrow_transaction_items (Detail)

| Field                 | Type        | Constraint | Keterangan          |
| --------------------- | ----------- | ---------- | ------------------- |
| id                    | uuid        | PK         | ID item             |
| borrow_transaction_id | uuid        | FK         | Relasi ke transaksi |
| tool_id               | uuid        | FK         | Alat yang dipinjam  |
| quantity              | int         | CHECK > 0  | Jumlah alat         |
| created_at            | timestamptz | NOT NULL   | Waktu dibuat        |

💡 **Catatan:**

- memungkinkan multi-item dalam 1 transaksi
- ini best practice (dipakai di e-commerce juga)

---

# 📤 return_transactions

| Field                 | Type        | Constraint | Keterangan          |
| --------------------- | ----------- | ---------- | ------------------- |
| id                    | uuid        | PK         | ID pengembalian     |
| borrow_transaction_id | uuid        | FK, UNIQUE | Relasi ke transaksi |
| returned_at           | timestamptz | NOT NULL   | Waktu dikembalikan  |
| late_days             | int         | CHECK ≥ 0  | Hari keterlambatan  |
| fine_amount           | numeric     | CHECK ≥ 0  | Total denda         |
| processed_by          | uuid        | FK         | Petugas             |
| created_at            | timestamptz | NOT NULL   | Waktu dicatat       |

💡 **Catatan penting:**

- 1 transaksi hanya bisa dikembalikan 1x (UNIQUE)
- `late_days` disimpan untuk efisiensi

---

# 📜 activity_logs

| Field       | Type        | Constraint | Keterangan                 |
| ----------- | ----------- | ---------- | -------------------------- |
| id          | uuid        | PK         | ID log                     |
| user_id     | uuid        | FK         | User pelaku                |
| action      | varchar     | NOT NULL   | Aksi (CREATE, UPDATE, dll) |
| entity      | varchar     | NOT NULL   | Tabel yang diubah          |
| entity_id   | uuid        | NULL       | ID data terkait            |
| description | text        | NULL       | Detail aksi                |
| created_at  | timestamptz | NOT NULL   | Waktu kejadian             |

💡 **Catatan:**

- penting untuk audit & debugging
- membantu tracking perubahan data

---

# 🔗 Relasi Antar Tabel

```text
users (1) → (N) borrow_transactions
users (1) → (N) activity_logs

categories (1) → (N) tools

borrow_transactions (1) → (N) borrow_transaction_items
tools (1) → (N) borrow_transaction_items

borrow_transactions (1) → (1) return_transactions
```

---

# 🔥 Insight Penting

## 1. Kenapa pakai header + detail?

→ supaya bisa:

- multi alat dalam 1 transaksi
- lebih fleksibel & scalable

---

## 2. Kenapa return dipisah?

→ karena:

- punya data sendiri (denda, tanggal)
- tidak semua transaksi langsung selesai

---

## 3. Kenapa pakai soft delete?

→ untuk:

- audit
- restore data
- keamanan

---

## 4. Kenapa pakai enum?

→ mencegah:

- typo
- data tidak valid
