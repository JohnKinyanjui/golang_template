-- name: GetUser :one
select
    u.id, u.full_name, u.email, u.phone,
    u.picture, u.status, u.google_uid,
    u.created_at, u.updated_at
from users u
where u.id = @user_id;


-- name: SignInWithGoogle :one
select id from users where email = @email::text and google_uid = @google_uid::text;

-- name: AuthenticateByEmail :one
select id from users where email = @email::text and password_hash = crypt(@password::text, password_hash) limit 1;

-- name: UpdateGoogleDetails :one
update users set email = @email::text , google_uid = @google_uid::text where id = @id returning id;

-- name: CreateUserWithEmail :one
insert into users (full_name, email, phone, password_hash)
values (
    @full_name::text,
    @email::text,
    nullif(@phone::text, ''),
    crypt(@password::text, gen_salt('bf'))
)
returning id;

-- name: UpdateMyAccount :one
update users
set
    full_name = case
        when sqlc.narg('full_name') is null or nullif(sqlc.narg('full_name')::text, '') is null then full_name
        else sqlc.narg('full_name')::text
    end,
    phone = case
        when sqlc.narg('phone') is null then phone
        else nullif(sqlc.narg('phone')::text, '')
    end
where id = @id
returning
    coalesce(picture, 'none') as picture,
    full_name,
    email,
    phone;
