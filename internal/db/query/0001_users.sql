-- name: GetUser :one
select picture, full_name, email from users where id = $1;

-- name: SignInWithGoogle :one
select id from users where email = @email::text and google_uid = @google_uid::text ;

-- name: UpdateGoogleDetails :one
update users set email = @email::text , google_uid = @google_uid::text where id = @id returning id;