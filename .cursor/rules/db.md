# SQL + sqlc Guidelines for Golang Backend

## 1. SQL Style & Structure

- Always write SQL in **lowercase**.
- Keep queries **short and concise** (`select ... limit 1` instead of complex joins unless required).
- use @ instead of $1, for example you can declare @name
- Always split database logic into separate files to keep schemas and queries clean and maintainable.
  - Use separate files for **major features or modules** (e.g., `002_table_a.up.sql`, `003_table_b.up.sql`).
  - Group **related tables** in the same file if they belong to the same feature (e.g., `table_a.sql` may also include `table_a_logs` or `table_a_metadata`).

## 2. Schema Migrations

- Store migrations in `./internal/db/schema/`.
- File format:
  - `0001_<table>.up.sql` → migration (create table).
  - ⚠️ `0001_users.up.sql` already exists — **do not adjust it** (you may add new columns if required).

## 3. Queries

- Store queries in `./internal/db/queries/`.
- File format: `0001_<table>.sql`.
- ⚠️ `0001_users.sql` already exists — **do not adjust the existing queries**.

## 4. Code Generation

- Use `sqlc.yaml` to link schema → queries → generated code.
- Generate Go code into `./internal/db/query/` with package name `query`.

## 5. Documentation Generation

- All documentation should be stored in `./docs`

## 6. Reducing Query Bloat

- **Use `*` for inserts/updates where possible.**  
  Example: instead of separate queries for partial updates, update all fields (or rely on Go to decide what to pass).

- **Rely on default queries.**  
  Only define queries that are actually used in your service layer.

  - Don’t write `update` if you only ever delete & recreate rows.
  - Don’t write `delete` if your app never deletes rows.

- **Limit each table to a minimal CRUD set.**  
  For most tables, define at most 3–4 queries:

  - `insert` (create)
  - `get by id` (read)
  - `list with pagination` (read multiple)
  - `delete by id` (delete)

- **Push update logic into Go.**  
  Instead of writing many `update_x_field` queries, do a single `update` query that updates multiple fields, or update in Go code where possible.

- **Group shared patterns.**  
  If multiple tables follow the same pattern (e.g., `created_at`, `updated_at`), keep queries consistent and avoid custom per-table variations.

## 7. Joins & JSON Objects

- **Use `json_build_object` for related tables.**  
  When fetching related data, wrap it into a JSON object instead of returning all columns flat.  
  This improves readability and avoids long `select` lists.

### Example

```sql
-- name: GetUser :one
select
  s.id,
  s.name,
  json_build_object(
    'id', sc.id,
    'name', sc.name,
    'address', sc.address
  ) as school
from table_users s
join address sc on s.address_id = sc.id
where s.id = $1
limit 1;
```

### 8. Casting

- For all queries (SELECT, INSERT, UPDATE, DELETE), cast `numeric` → `::int4` or `::float8`, and `timestamp` → `::timestamptz`, so sqlc generates Go-native types (`int32`, `float64`, `time.Time`) instead of pgtype wrappers.

### 9. Queries

- Instead, write the query to return exactly the shape you need (e.g. JSON, nested object, renamed fields), so your service can just forward the result.

### 10. Multitenancy Scenario

- Enforce multi-tenancy by linking each `user_id` to a tenant ID (e.g., `organization_id`, `store_id`, `school_id`).
- For related domain tables (e.g., events, parents, storage), always store and query by the tenant ID instead of `user_id`.
- On each request, resolve the tenant ID from `ctx.Get("user_id").(string)` and use it for filtering queries.
- Add a field in the `users` table to track the current active tenant (e.g., `current_organization_id uuid` referencing `organizations(id)` or similar).
- Provide a query (e.g., `GetCurrentUserOrg`) in `./internal/db/queries/` that returns the current active tenant ID for a given user.
