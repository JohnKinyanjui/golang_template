# Backend Development Guidelines

## 1. Project Structure

- Business logic lives in `./internal/services/`, grouped by feature (`accounts/`, `auth/`, `storage/`).
- Each service contains:
  - `*_service.go` → Business logic functions
  - `data.go` → DTOs + validation
  - `utils.go` → helpers (optional)
- API routers live in `./internal/api/` (`router.go` + feature routers).
- Handlers in ./internal/api/{service}/\* must be split by method into get.go, post.go, put.go, and delete.go.
- Handlers must return errors in an if err != nil block with api_helpers.ResultSimple(c, err.Understandable, err.Err), and only return success with a nil error.

## 2. Service Functions

- Use descriptive names (`CreateUser`, `GetUserByID`).
- Return `(result, *helpers.ResultError)` or `*helpers.ResultError`.
- Each function should follow **single responsibility**.
- services error function should be \*helpers.ResultError

## 3. Data & Validation

```go
// Example in data.go
type CreateUserParams struct {
    FullName string `json:"full_name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
}
func (p *CreateUserParams) validate() error {
    return nil // add validation logic
}
```

## 4. Database

- We use **sqlc with pgx** for database access.
- All generated code must use **pgtype.\*** types (✅) and never `sql.Null*`.
- Ensure `sqlc.yaml` is configured with:
  ```yaml
  sql_package: "pgx/v5"
  ```
- Single queries → db.Query.Method(ctx, query.Params{...})
- Transactions → qtx := db.Query.WithTx(tx) then call queries on qtx

## 5. No More Reshappping

Don’t manually transform or reshape query results in Go.

## 6. Getting User ID

- Always retrieve the authenticated user ID from context using `id := ctx.Get("user_id").(string)`.
- This should be used whenever a request is tied to the current profile or in multi-tenant scenarios, where the user’s ID replaces an explicit `user_id` parameter.
- note user_id is a string in most cases you will need to convert it into a uuid to work i recommend uuid.MustParse() since it by default uuid

## 7, Multitenancy

- Handlers and services should always call a query to resolve the active tenant (e.g., `organization_id`, `store_id`, `school_id`, etc.) for the current user, instead of accepting these IDs from the request.

## 8. Folder Structure

Always organize features into separate folders (e.g., `feature_a/`, `feature_b/`) to keep the project clean, maintainable, and scalable.
