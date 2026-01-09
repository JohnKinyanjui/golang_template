# Profile / Account API

All endpoints require `Authorization: Bearer <jwt>` header.

---

## Get My Profile

**GET** `/api/v1/account/my`

Fetch the current authenticated user's profile.

### Response (200 OK) — Get My Profile

```json
{
  "picture": "url_or_none",
  "username": "string",
  "email": "string"
}
```

---

## Update My Profile

**PUT** `/api/v1/account/my`

Update profile fields for the current user. Only include fields you want to change.

### Request Body

```json
{
  "username": "string",
  "about": "string",
  "location": "string",
  "looking_for": "string"
}
```

### Response (200 OK) — Update My Profile

```json
{
  "picture": "url_or_none",
  "username": "string",
  "email": "string",
  "about": "string",
  "location": "string",
  "looking_for": "string"
}
```

---

## Upload Profile Picture

**POST** `/api/v1/account/picture`

Upload or replace the profile picture. Expect `multipart/form-data` with a file field named `picture`.

### Request Body (multipart/form-data)

- `picture`: Image file (jpeg/png/webp)

### Response (200 OK) — Upload Profile Picture

```json
{
  "picture": "url_to_uploaded_image"
}
```

---

## Delete Account

**DELETE** `/api/v1/account/my`

Deactivate/delete the current user's account.

### Response (200 OK) — Delete Account

```json
{
  "message": "account deleted"
}
```

---

### Notes

- All endpoints return JWT-based authentication errors as `{ "error": "..." }`.
- Profile picture URLs are stored as strings; `"none"` indicates no picture set.
- `username` is unique and must be provided on update if changed.
- Optional fields (`about`, `location`, `looking_for`) can be omitted or set to `null` to clear them.