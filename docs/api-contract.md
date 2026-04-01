# 📋 API Contract – SIPALA

> **Version:** v1  
> **Base URL:** `/api`  
> **Content-Type:** `application/json`

---

## 🔐 Response Format Standard

Semua response API mengikuti format berikut:

```json
{
  "code": 200,
  "message": "Success",
  "result": { ... }
}
```

### HTTP Status Codes

| Code | Meaning               |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 409  | Conflict              |
| 500  | Internal Server Error |

---

# 🔑 Authentication Service

**Base Path:** `/api/auth/v1`

---

## http-login

| Attribute       | Value                                                                                                                                                                                           |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                          |
| **Path**        | `/api/auth/v1/login`                                                                                                                                                                            |
| **Description** | Melakukan autentikasi user dengan email dan password. Jika valid, generate access token (JWT) dan refresh token. Simpan refresh token ke database/session. Return token pair beserta data user. |
| **Auth**        | `none`                                                                                                                                                                                          |

### Request Body

```json
{
  "email": "string (required)",
  "password": "string (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Login successful",
  "result": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "user": {
      "id": "uuid",
      "name": "string",
      "email": "string",
      "role": "admin|staff|borrower"
    }
  }
}
```

---

## http-logout

| Attribute       | Value                                                                                                                    |
| --------------- | ------------------------------------------------------------------------------------------------------------------------ |
| **Method**      | `POST`                                                                                                                   |
| **Path**        | `/api/auth/v1/logout`                                                                                                    |
| **Description** | Melakukan invalidasi token. Hapus/blacklist refresh token dari database. Client juga harus menghapus token dari storage. |
| **Auth**        | `Bearer` (any role)                                                                                                      |

### Request Body

```json
{
  "refresh_token": "string (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Logout successful",
  "result": null
}
```

---

## http-register

| Attribute       | Value                                                                                                                                                                                                    |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                   |
| **Path**        | `/api/auth/v1/register`                                                                                                                                                                                  |
| **Description** | Mendaftarkan user baru dengan role default `borrower`. Validasi email unique, hash password dengan bcrypt/argon2, simpan ke tabel users. Tidak langsung login, user harus login manual setelah register. |
| **Auth**        | `none`                                                                                                                                                                                                   |

### Request Body

```json
{
  "name": "string (required)",
  "email": "string (required, unique)",
  "password": "string (required, min 6 chars)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "User registered successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "role": "borrower",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-refresh-token

| Attribute       | Value                                                                                                                                                                                   |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                  |
| **Path**        | `/api/auth/v1/refresh-token`                                                                                                                                                            |
| **Description** | Generate access token baru menggunakan refresh token yang valid. Validasi refresh token dari database, cek expired, generate token pair baru, invalidate token lama, simpan token baru. |
| **Auth**        | `none` (tapi butuh valid refresh token)                                                                                                                                                 |

### Request Body

```json
{
  "refresh_token": "string (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Token refreshed successfully",
  "result": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "bmV3IHJlZnJlc2ggdG9rZW4...",
    "token_type": "Bearer",
    "expires_in": 3600
  }
}
```

---

## http-get-profile

| Attribute       | Value                                                                                                                                     |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                     |
| **Path**        | `/api/auth/v1/profile`                                                                                                                    |
| **Description** | Mengambil data profile user yang sedang login berdasarkan token. Decode JWT, ambil user_id, query ke tabel users (exclude password_hash). |
| **Auth**        | `Bearer` (any role)                                                                                                                       |

### Request Query

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "role": "admin|staff|borrower",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-change-password

| Attribute       | Value                                                                                                                                                                 |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `PUT`                                                                                                                                                                 |
| **Path**        | `/api/auth/v1/change-password`                                                                                                                                        |
| **Description** | Mengubah password user yang sedang login. Validasi old_password match dengan hash di database, validasi new_password strength, hash new_password, update ke database. |
| **Auth**        | `Bearer` (any role)                                                                                                                                                   |

### Request Body

```json
{
  "old_password": "string (required)",
  "new_password": "string (required, min 6 chars)",
  "confirm_password": "string (required, must match new_password)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Password changed successfully",
  "result": null
}
```

---

# 👤 Users Service

**Base Path:** `/api/users/v1`

---

## http-create-user

| Attribute       | Value                                                                                                                                                         |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                        |
| **Path**        | `/api/users/v1/users`                                                                                                                                         |
| **Description** | Membuat user baru oleh admin/staff. Validasi email unique, hash password, simpan ke tabel users. Role bisa ditentukan (default borrower). Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                       |

### Request Body

```json
{
  "name": "string (required)",
  "email": "string (required, unique)",
  "password": "string (required, min 6 chars)",
  "role": "admin|staff|borrower (default: borrower)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "User created successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "role": "string",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-update-user

| Attribute       | Value                                                                                                                                                                                                          |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `PUT`                                                                                                                                                                                                          |
| **Path**        | `/api/users/v1/users/{id}`                                                                                                                                                                                     |
| **Description** | Update data user. Admin bisa update semua field termasuk role. User biasa hanya bisa update profile sendiri (tidak bisa ganti role). Validasi email unique jika diubah. Update updated_at. Catat activity log. |
| **Auth**        | `Bearer` (admin, staff, owner)                                                                                                                                                                                 |

### Request Body

```json
{
  "name": "string (optional)",
  "email": "string (optional, unique)",
  "role": "admin|staff|borrower (optional, admin only)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "User updated successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "role": "string",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-delete-user

| Attribute       | Value                                                                                                                                                      |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `DELETE`                                                                                                                                                   |
| **Path**        | `/api/users/v1/users/{id}`                                                                                                                                 |
| **Description** | Soft delete user dengan mengisi deleted_at. Cek apakah user punya transaksi aktif (borrow yang belum returned), jika ada tolak delete. Catat activity log. |
| **Auth**        | `Bearer` (admin only)                                                                                                                                      |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "User deleted successfully",
  "result": {
    "id": "uuid",
    "deleted_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-user-by-id

| Attribute       | Value                                                                                                                                 |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                 |
| **Path**        | `/api/users/v1/users/{id}`                                                                                                            |
| **Description** | Ambil detail user by ID. Exclude deleted users (kecuali admin dengan param include_deleted). Return data lengkap tanpa password_hash. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                               |

### Request Query

```json
{
  "include_deleted": "boolean (optional, default: false, admin only)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "role": "string",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "deleted_at": "null|timestamp"
  }
}
```

---

## http-get-list-users

| Attribute       | Value                                                                                                                                                     |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                     |
| **Path**        | `/api/users/v1/users`                                                                                                                                     |
| **Description** | List users dengan pagination, search, dan filter. Support filter by role, search by name/email. Exclude soft deleted by default. Sort by created_at desc. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                   |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "search": "string (optional, search name/email)",
  "role": "admin|staff|borrower (optional)",
  "include_deleted": "boolean (optional, default: false, admin only)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "name": "string",
        "email": "string",
        "role": "string",
        "created_at": "2024-01-01T00:00:00Z"
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 🏷️ Categories Service

**Base Path:** `/api/categories/v1`

---

## http-create-category

| Attribute       | Value                                                                                             |
| --------------- | ------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                            |
| **Path**        | `/api/categories/v1/categories`                                                                   |
| **Description** | Membuat kategori alat baru. Validasi nama unique, simpan ke tabel categories. Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                                           |

### Request Body

```json
{
  "name": "string (required, unique)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "Category created successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-update-category

| Attribute       | Value                                                                              |
| --------------- | ---------------------------------------------------------------------------------- |
| **Method**      | `PUT`                                                                              |
| **Path**        | `/api/categories/v1/categories/{id}`                                               |
| **Description** | Update nama kategori. Validasi nama unique, update updated_at. Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                            |

### Request Body

```json
{
  "name": "string (required, unique)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Category updated successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-delete-category

| Attribute       | Value                                                                                                                                                                            |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `DELETE`                                                                                                                                                                         |
| **Path**        | `/api/categories/v1/categories/{id}`                                                                                                                                             |
| **Description** | Hapus kategori. Cek apakah ada tools yang menggunakan kategori ini, jika ada tolak delete atau pindahkan ke kategori default. Soft delete dengan deleted_at. Catat activity log. |
| **Auth**        | `Bearer` (admin only)                                                                                                                                                            |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Category deleted successfully",
  "result": {
    "id": "uuid",
    "deleted_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-category-by-id

| Attribute       | Value                                                                     |
| --------------- | ------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                     |
| **Path**        | `/api/categories/v1/categories/{id}`                                      |
| **Description** | Ambil detail kategori by ID. Include count tools dalam kategori tersebut. |
| **Auth**        | `Bearer` (any role)                                                       |

### Request Query

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "name": "string",
    "tools_count": 5,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-list-categories

| Attribute       | Value                                                               |
| --------------- | ------------------------------------------------------------------- |
| **Method**      | `GET`                                                               |
| **Path**        | `/api/categories/v1/categories`                                     |
| **Description** | List semua kategori dengan pagination dan search. Sort by name asc. |
| **Auth**        | `Bearer` (any role)                                                 |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "search": "string (optional, search by name)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "name": "string",
        "tools_count": 5
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 🧰 Tools Service

**Base Path:** `/api/tools/v1`

---

## http-create-tool

| Attribute       | Value                                                                                                  |
| --------------- | ------------------------------------------------------------------------------------------------------ |
| **Method**      | `POST`                                                                                                 |
| **Path**        | `/api/tools/v1/tools`                                                                                  |
| **Description** | Membuat alat baru. Validasi category_id exists, stock >= 0, simpan ke tabel tools. Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                                                |

### Request Body

```json
{
  "name": "string (required)",
  "category_id": "uuid (required)",
  "stock": "integer (required, min: 0)",
  "description": "string (optional)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "Tool created successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "category_id": "uuid",
    "category_name": "string",
    "stock": 10,
    "description": "string",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-update-tool

| Attribute       | Value                                                                                                                        |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `PUT`                                                                                                                        |
| **Path**        | `/api/tools/v1/tools/{id}`                                                                                                   |
| **Description** | Update data alat. Validasi category_id exists jika diubah, stock tidak boleh negatif. Update updated_at. Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                      |

### Request Body

```json
{
  "name": "string (optional)",
  "category_id": "uuid (optional)",
  "stock": "integer (optional, min: 0)",
  "description": "string (optional)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Tool updated successfully",
  "result": {
    "id": "uuid",
    "name": "string",
    "category_id": "uuid",
    "category_name": "string",
    "stock": 15,
    "description": "string",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-delete-tool

| Attribute       | Value                                                                                                                                                                |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `DELETE`                                                                                                                                                             |
| **Path**        | `/api/tools/v1/tools/{id}`                                                                                                                                           |
| **Description** | Soft delete alat. Cek apakah ada borrow_transaction_items yang statusnya bukan returned (masih dipinjam), jika ada tolak delete. Set deleted_at. Catat activity log. |
| **Auth**        | `Bearer` (admin only)                                                                                                                                                |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Tool deleted successfully",
  "result": {
    "id": "uuid",
    "deleted_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-tool-by-id

| Attribute       | Value                                                                                                         |
| --------------- | ------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                         |
| **Path**        | `/api/tools/v1/tools/{id}`                                                                                    |
| **Description** | Ambil detail alat by ID. Include data kategori, hitung available_stock (stock - jumlah yang sedang dipinjam). |
| **Auth**        | `Bearer` (any role)                                                                                           |

### Request Query

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "name": "string",
    "category": {
      "id": "uuid",
      "name": "string"
    },
    "stock": 10,
    "available_stock": 7,
    "borrowed_count": 3,
    "description": "string",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-list-tools

| Attribute       | Value                                                                                                                                                           |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                           |
| **Path**        | `/api/tools/v1/tools`                                                                                                                                           |
| **Description** | List alat dengan pagination, search, dan filter. Support filter by category_id, search by name, filter by availability (available_stock > 0). Sort by name asc. |
| **Auth**        | `Bearer` (any role)                                                                                                                                             |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "search": "string (optional, search by name)",
  "category_id": "uuid (optional)",
  "available_only": "boolean (optional, default: false)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "name": "string",
        "category": {
          "id": "uuid",
          "name": "string"
        },
        "stock": 10,
        "available_stock": 7,
        "description": "string"
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 📥 Borrow Service

**Base Path:** `/api/borrows/v1`

---

## http-create-borrow-transaction

| Attribute       | Value                                                                                                                                                                                                               |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                              |
| **Path**        | `/api/borrows/v1/borrows`                                                                                                                                                                                           |
| **Description** | Membuat header transaksi peminjaman (draft). Set borrower_id dari token, status default `pending`, borrow_date default now(), due_date wajib diisi (harus >= borrow_date). Return transaction ID untuk nambah item. |
| **Auth**        | `Bearer` (borrower, admin, staff)                                                                                                                                                                                   |

### Request Body

```json
{
  "due_date": "timestamp (required, ISO 8601, >= today)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "Borrow transaction created",
  "result": {
    "id": "uuid",
    "borrower_id": "uuid",
    "status": "pending",
    "borrow_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-01-08T00:00:00Z",
    "items": [],
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-add-borrow-item

| Attribute       | Value                                                                                                                                                                                                                                                      |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                                                                     |
| **Path**        | `/api/borrows/v1/borrows/{borrow_id}/items`                                                                                                                                                                                                                |
| **Description** | Menambahkan item ke transaksi peminjaman. Validasi borrow_id exists dan status masih `pending`, validasi tool_id exists dan not soft deleted, cek available_stock >= quantity, insert ke borrow_transaction_items. Update/lock stock sementara (optional). |
| **Auth**        | `Bearer` (borrower - owner only, admin, staff)                                                                                                                                                                                                             |

### Request Body

```json
{
  "tool_id": "uuid (required)",
  "quantity": "integer (required, min: 1)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "Item added to borrow transaction",
  "result": {
    "id": "uuid",
    "borrow_transaction_id": "uuid",
    "tool_id": "uuid",
    "tool_name": "string",
    "quantity": 2,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-remove-borrow-item

| Attribute       | Value                                                                                                                                                                                                          |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `DELETE`                                                                                                                                                                                                       |
| **Path**        | `/api/borrows/v1/borrows/{borrow_id}/items/{item_id}`                                                                                                                                                          |
| **Description** | Menghapus item dari transaksi peminjaman. Validasi borrow_id exists dan status `pending`, validasi item_id exists dalam transaksi tersebut, delete dari borrow_transaction_items. Release stock lock jika ada. |
| **Auth**        | `Bearer` (borrower - owner only, admin, staff)                                                                                                                                                                 |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Item removed from borrow transaction",
  "result": {
    "removed_item_id": "uuid"
  }
}
```

---

## http-submit-borrow

| Attribute       | Value                                                                                                                                                                                                                                                        |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Method**      | `POST`                                                                                                                                                                                                                                                       |
| **Path**        | `/api/borrows/v1/borrows/{borrow_id}/submit`                                                                                                                                                                                                                 |
| **Description** | Submit transaksi peminjaman untuk approval. Validasi borrow_id exists dan status `pending`, cek minimal ada 1 item, validasi ulang stock availability untuk semua items, update status tetap `pending` (menunggu approval). Kirim notifikasi ke staff/admin. |
| **Auth**        | `Bearer` (borrower - owner only, admin, staff)                                                                                                                                                                                                               |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Borrow transaction submitted for approval",
  "result": {
    "id": "uuid",
    "status": "pending",
    "submitted_at": "2024-01-01T00:00:00Z",
    "total_items": 3,
    "total_quantity": 5
  }
}
```

---

## http-get-borrow-by-id

| Attribute       | Value                                                                                                                                                                          |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Method**      | `GET`                                                                                                                                                                          |
| **Path**        | `/api/borrows/v1/borrows/{id}`                                                                                                                                                 |
| **Description** | Ambil detail transaksi peminjaman by ID. Include semua items (dengan detail tool), borrower info, approval info jika sudah diapprove. Cek authorization (owner, admin, staff). |
| **Auth**        | `Bearer` (borrower - owner only, admin, staff)                                                                                                                                 |

### Request Query

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "borrower": {
      "id": "uuid",
      "name": "string",
      "email": "string"
    },
    "status": "pending|approved|rejected|returned",
    "borrow_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-01-08T00:00:00Z",
    "items": [
      {
        "id": "uuid",
        "tool_id": "uuid",
        "tool_name": "string",
        "quantity": 2
      }
    ],
    "approved_by": "uuid|null",
    "approved_at": "timestamp|null",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## http-get-list-borrows

| Attribute       | Value                                                                                                                                                                 |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                                 |
| **Path**        | `/api/borrows/v1/borrows`                                                                                                                                             |
| **Description** | List semua transaksi peminjaman dengan pagination, filter, dan search. Support filter by status, borrower_id, date range. Sort by created_at desc. Untuk admin/staff. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                               |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "status": "pending|approved|rejected|returned (optional)",
  "borrower_id": "uuid (optional)",
  "start_date": "date (optional, ISO 8601)",
  "end_date": "date (optional, ISO 8601)",
  "search": "string (optional, search borrower name)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "borrower": {
          "id": "uuid",
          "name": "string"
        },
        "status": "string",
        "borrow_date": "2024-01-01T00:00:00Z",
        "due_date": "2024-01-08T00:00:00Z",
        "total_items": 3,
        "total_quantity": 5
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

## http-get-my-borrows

| Attribute       | Value                                                                                                                                                             |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                             |
| **Path**        | `/api/borrows/v1/my-borrows`                                                                                                                                      |
| **Description** | List transaksi peminjaman milik user yang sedang login. Filter otomatis by borrower_id dari token. Support filter by status, pagination. Sort by created_at desc. |
| **Auth**        | `Bearer` (any role)                                                                                                                                               |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "status": "pending|approved|rejected|returned (optional)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": {
      "data": [
        {
          "id": "uuid",
          "status": "string",
          "borrow_date": "2024-01-01T00:00:00Z",
          "due_date": "2024-01-08T00:00:00Z",
          "total_items": 3,
          "total_quantity": 5,
          "is_overdue": false
        }
      ],
      "metadata": {
        "current_page": "int",
        "page_size": "int",
        "total_pages": "int",
        "total_records": "int",
        "has_next": "boolean",
        "has_prev": "boolean"
      }
    }
  }
}
```

---

# ✅ Borrow Approval Service

**Base Path:** `/api/approvals/v1`

---

## http-approve-borrow

| Attribute       | Value                                                                                                                                                                                                                                                                                  |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                                                                                                 |
| **Path**        | `/api/approvals/v1/borrows/{borrow_id}/approve`                                                                                                                                                                                                                                        |
| **Description** | Menyetujui transaksi peminjaman. Validasi borrow_id exists dan status `pending`, set status jadi `approved`, set approved_by dari token, set approved_at now(), kurangi stock tools sesuai quantity yang dipinjam (decrement stock). Catat activity log. Kirim notifikasi ke borrower. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                                                                                                                |

### Request Body

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Borrow transaction approved",
  "result": {
    "id": "uuid",
    "status": "approved",
    "approved_by": {
      "id": "uuid",
      "name": "string"
    },
    "approved_at": "2024-01-01T00:00:00Z",
    "borrow_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-01-08T00:00:00Z"
  }
}
```

---

## http-reject-borrow

| Attribute       | Value                                                                                                                                                                                                                                                                                           |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                                                                                                          |
| **Path**        | `/api/approvals/v1/borrows/{borrow_id}/reject`                                                                                                                                                                                                                                                  |
| **Description** | Menolak transaksi peminjaman. Validasi borrow_id exists dan status `pending`, set status jadi `rejected`, set approved_by dari token, set approved_at now(), release stock lock jika ada (kembalikan stock). Optional: bisa tambahkan reason. Catat activity log. Kirim notifikasi ke borrower. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                                                                                                                         |

### Request Body

```json
{
  "reason": "string (optional)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Borrow transaction rejected",
  "result": {
    "id": "uuid",
    "status": "rejected",
    "rejected_by": {
      "id": "uuid",
      "name": "string"
    },
    "rejected_at": "2024-01-01T00:00:00Z",
    "reason": "string|null"
  }
}
```

---

# 📤 Return Service

**Base Path:** `/api/returns/v1`

---

## http-create-return

| Attribute       | Value                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                                                                                                                                                                                                                                      |
| **Path**        | `/api/returns/v1/returns`                                                                                                                                                                                                                                                                                                                                                                                                   |
| **Description** | Membuat transaksi pengembalian. Validasi borrow_transaction_id exists dan status `approved`, cek belum ada return transaction (UNIQUE), set returned_at now(), hitung late_days (returned_at - due_date, min 0), hitung fine_amount berdasarkan late_days \* rate per hari, insert ke return_transactions. Update borrow_transactions status jadi `returned`. Kembalikan stock tools (increment stock). Catat activity log. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                                                                                                                                                                                                                                                     |

### Request Body

```json
{
  "borrow_transaction_id": "uuid (required)"
}
```

### Response 201

```json
{
  "code": 201,
  "message": "Return transaction created",
  "result": {
    "id": "uuid",
    "borrow_transaction_id": "uuid",
    "returned_at": "2024-01-10T00:00:00Z",
    "late_days": 2,
    "fine_amount": 20000,
    "processed_by": {
      "id": "uuid",
      "name": "string"
    },
    "borrow_details": {
      "borrower_name": "string",
      "borrow_date": "2024-01-01T00:00:00Z",
      "due_date": "2024-01-08T00:00:00Z",
      "items": [
        {
          "tool_name": "string",
          "quantity": 2
        }
      ]
    }
  }
}
```

---

## http-calculate-fine

| Attribute       | Value                                                                                                                                                                                                                       |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `POST`                                                                                                                                                                                                                      |
| **Path**        | `/api/returns/v1/calculate-fine`                                                                                                                                                                                            |
| **Description** | Kalkulasi denda tanpa membuat transaksi (preview). Validasi borrow_transaction_id exists dan status `approved`, hitung late_days dari due_date sampai now(), hitung fine_amount. Return kalkulasi tanpa insert ke database. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                                                     |

### Request Body

```json
{
  "borrow_transaction_id": "uuid (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Fine calculated",
  "result": {
    "borrow_transaction_id": "uuid",
    "due_date": "2024-01-08T00:00:00Z",
    "current_date": "2024-01-10T00:00:00Z",
    "late_days": 2,
    "fine_per_day": 10000,
    "fine_amount": 20000,
    "currency": "IDR"
  }
}
```

---

## http-get-return-by-id

| Attribute       | Value                                                                                                                                     |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                     |
| **Path**        | `/api/returns/v1/returns/{id}`                                                                                                            |
| **Description** | Ambil detail transaksi pengembalian by ID. Include borrow transaction details, borrower info, processed_by info, items yang dikembalikan. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                   |

### Request Query

`none`

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "id": "uuid",
    "borrow_transaction": {
      "id": "uuid",
      "borrower": {
        "id": "uuid",
        "name": "string",
        "email": "string"
      },
      "borrow_date": "2024-01-01T00:00:00Z",
      "due_date": "2024-01-08T00:00:00Z"
    },
    "returned_at": "2024-01-10T00:00:00Z",
    "late_days": 2,
    "fine_amount": 20000,
    "processed_by": {
      "id": "uuid",
      "name": "string"
    },
    "items": [
      {
        "tool_name": "string",
        "quantity": 2
      }
    ]
  }
}
```

---

## http-get-list-returns

| Attribute       | Value                                                                                                                                                                    |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Method**      | `GET`                                                                                                                                                                    |
| **Path**        | `/api/returns/v1/returns`                                                                                                                                                |
| **Description** | List transaksi pengembalian dengan pagination dan filter. Support filter by date range, borrower_id, filter by ada/tidak ada denda (has_fine). Sort by returned_at desc. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                  |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "borrower_id": "uuid (optional)",
  "start_date": "date (optional)",
  "end_date": "date (optional)",
  "has_fine": "boolean (optional)",
  "search": "string (optional, search borrower name)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "borrow_transaction_id": "uuid",
        "borrower_name": "string",
        "returned_at": "2024-01-10T00:00:00Z",
        "late_days": 2,
        "fine_amount": 20000
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 📊 Monitoring Service

**Base Path:** `/api/monitoring/v1`

---

## http-get-active-borrows

| Attribute       | Value                                                                                                                                                                                                     |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                                                                     |
| **Path**        | `/api/monitoring/v1/active-borrows`                                                                                                                                                                       |
| **Description** | List peminjaman yang sedang aktif (status `approved` dan belum returned). Include sisa hari sampai due_date, flag overdue jika sudah lewat due_date. Sort by due_date asc (yang paling mendesak di atas). |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                                   |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "search": "string (optional, search borrower name)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "borrower": {
          "id": "uuid",
          "name": "string",
          "email": "string"
        },
        "borrow_date": "2024-01-01T00:00:00Z",
        "due_date": "2024-01-08T00:00:00Z",
        "days_remaining": 2,
        "is_overdue": false,
        "total_items": 3,
        "items": [
          {
            "tool_name": "string",
            "quantity": 2
          }
        ]
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

## http-get-overdue-borrows

| Attribute       | Value                                                                                                                                                                                              |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                                                              |
| **Path**        | `/api/monitoring/v1/overdue-borrows`                                                                                                                                                               |
| **Description** | List peminjaman yang sudah overdue (status `approved` dan due_date < now()). Hitung late_days, estimasi denda jika dikembalikan hari ini. Sort by due_date asc (yang paling lama overdue di atas). |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                                                            |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "search": "string (optional, search borrower name)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "borrower": {
          "id": "uuid",
          "name": "string",
          "email": "string"
        },
        "borrow_date": "2024-01-01T00:00:00Z",
        "due_date": "2024-01-05T00:00:00Z",
        "late_days": 5,
        "estimated_fine": 50000,
        "total_items": 3,
        "items": [
          {
            "tool_name": "string",
            "quantity": 2
          }
        ]
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 📈 Reporting Service

**Base Path:** `/api/reports/v1`

---

## http-get-borrow-report

| Attribute       | Value                                                                                                                                                      |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                      |
| **Path**        | `/api/reports/v1/borrows`                                                                                                                                  |
| **Description** | Report statistik peminjaman. Aggregasi data peminjaman by range tanggal. Include total borrow, breakdown by status, most borrowed tools, peak borrow days. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                                    |

### Request Query

```json
{
  "start_date": "date (required)",
  "end_date": "date (required)",
  "group_by": "day|week|month (optional, default: day)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "period": {
      "start_date": "2024-01-01",
      "end_date": "2024-01-31"
    },
    "summary": {
      "total_transactions": 150,
      "total_items_borrowed": 320,
      "approved": 140,
      "rejected": 10,
      "pending": 0
    },
    "top_tools": [
      {
        "tool_id": "uuid",
        "tool_name": "string",
        "borrow_count": 50,
        "total_quantity": 100
      }
    ],
    "by_date": [
      {
        "date": "2024-01-01",
        "transaction_count": 5,
        "item_count": 12
      }
    ]
  }
}
```

---

## http-get-return-report

| Attribute       | Value                                                                                                                                            |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Method**      | `GET`                                                                                                                                            |
| **Path**        | `/api/reports/v1/returns`                                                                                                                        |
| **Description** | Report statistik pengembalian. Aggregasi data pengembalian by range tanggal. Include total return, on-time vs late returns, average return time. |
| **Auth**        | `Bearer` (admin, staff)                                                                                                                          |

### Request Query

```json
{
  "start_date": "date (required)",
  "end_date": "date (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "period": {
      "start_date": "2024-01-01",
      "end_date": "2024-01-31"
    },
    "summary": {
      "total_returns": 140,
      "on_time": 120,
      "late": 20,
      "on_time_rate": 85.7
    },
    "average_late_days": 3.5,
    "by_date": [
      {
        "date": "2024-01-10",
        "return_count": 8,
        "late_count": 2
      }
    ]
  }
}
```

---

## http-get-fine-report

| Attribute       | Value                                                                                                             |
| --------------- | ----------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                             |
| **Path**        | `/api/reports/v1/fines`                                                                                           |
| **Description** | Report statistik denda. Aggregasi total denda by range tanggal. Breakdown by borrower, by tool, trends over time. |
| **Auth**        | `Bearer` (admin, staff)                                                                                           |

### Request Query

```json
{
  "start_date": "date (required)",
  "end_date": "date (required)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "period": {
      "start_date": "2024-01-01",
      "end_date": "2024-01-31"
    },
    "summary": {
      "total_fine_amount": 500000,
      "total_transactions_with_fine": 20,
      "average_fine": 25000,
      "max_fine": 100000
    },
    "top_borrowers_with_fine": [
      {
        "borrower_id": "uuid",
        "borrower_name": "string",
        "fine_count": 3,
        "total_fine": 75000
      }
    ],
    "by_date": [
      {
        "date": "2024-01-10",
        "fine_count": 2,
        "fine_amount": 40000
      }
    ]
  }
}
```

---

# 📜 Logs Service

**Base Path:** `/api/logs/v1`

---

## http-get-logs

| Attribute       | Value                                                                                                                                                             |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Method**      | `GET`                                                                                                                                                             |
| **Path**        | `/api/logs/v1/logs`                                                                                                                                               |
| **Description** | List activity logs dengan pagination dan filter. Support filter by user_id, action (CREATE, UPDATE, DELETE), entity (tabel), date range. Sort by created_at desc. |
| **Auth**        | `Bearer` (admin only)                                                                                                                                             |

### Request Query

```json
{
  "page": "integer (optional, default: 1)",
  "limit": "integer (optional, default: 10, max: 100)",
  "user_id": "uuid (optional)",
  "action": "CREATE|UPDATE|DELETE|LOGIN|LOGOUT (optional)",
  "entity": "users|tools|categories|borrow_transactions|return_transactions (optional)",
  "start_date": "date (optional)",
  "end_date": "date (optional)"
}
```

### Response 200

```json
{
  "code": 200,
  "message": "Success",
  "result": {
    "data": [
      {
        "id": "uuid",
        "user": {
          "id": "uuid",
          "name": "string"
        },
        "action": "CREATE",
        "entity": "tools",
        "entity_id": "uuid",
        "description": "Created new tool: Hammer",
        "created_at": "2024-01-01T10:00:00Z"
      }
    ],
    "metadata": {
      "current_page": "int",
      "page_size": "int",
      "total_pages": "int",
      "total_records": "int",
      "has_next": "boolean",
      "has_prev": "boolean"
    }
  }
}
```

---

# 📝 Summary

## Endpoint Count by Service

| Service         | Endpoints |
| --------------- | --------- |
| Authentication  | 6         |
| Users           | 5         |
| Categories      | 5         |
| Tools           | 5         |
| Borrow          | 7         |
| Borrow Approval | 2         |
| Return          | 4         |
| Monitoring      | 2         |
| Reporting       | 3         |
| Logs            | 1         |
| **Total**       | **40**    |

## Auth Matrix

| Role         | Permissions                                                                                                                                                                                 |
| ------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Admin**    | Full access (all endpoints)                                                                                                                                                                 |
| **Staff**    | All except: delete-user, delete-category, delete-tool, get-logs                                                                                                                             |
| **Borrower** | login, logout, register, refresh-token, get-profile, change-password, create-borrow-transaction, add-borrow-item, remove-borrow-item, submit-borrow, get-borrow-by-id (own), get-my-borrows |
