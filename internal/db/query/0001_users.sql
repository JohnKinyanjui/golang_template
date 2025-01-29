-- name: GetUser :one
select picture, full_name, email from users where id = $1;