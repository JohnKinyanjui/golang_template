# Auth API

## Register with Email

**POST** `/api/v1/auth/register/email`

Register a new user with email and password.

### Request Body
```json
{
  "username": "string",
  "email": "string",
  "phone": "string",
  "password": "string",
  "gender": "male|female|other",
  "about": "string",
  "location": "string",
  "looking_for": "string"
}
```

### Response (200 OK)
```json
{
  "token": "jwt_token_string"
}
```

---

## Login with Email

**POST** `/api/v1/auth/email`

Login with email and password.

### Request Body
```json
{
  "email": "string",
  "password": "string"
}
```

### Response (200 OK)
```json
{
  "token": "jwt_token_string"
}
```

---

### Notes
- All endpoints return JWT on success.
- JWT should be sent in `Authorization: Bearer <token>` header for protected routes.
- Passwords are hashed before storage.