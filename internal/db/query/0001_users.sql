-- name: GetUser :one
select 
    u.id,
    u.username,
    u.email,
    u.phone,
    u.picture,
    u.gender,
    u.is_verified,
    u.last_login_at,
    u.status,
    u.created_at,
    u.updated_at
from users u
where u.id = @user_id;

-- name: SignInWithGoogle :one
select id from users where email = @email::text and google_uid = @google_uid::text;

-- name: AuthenticateByEmail :one
select id from users where email = @email::text and password_hash = crypt(@password::text, password_hash) limit 1;

-- name: UpdateGoogleDetails :one
update users set email = @email::text , google_uid = @google_uid::text where id = @id returning id;

-- name: CreateUserWithEmail :one
insert into users (username, email, phone, password_hash, gender, about, location, looking_for)
values (
    @username::text,
    @email::text,
    nullif(@phone::text, ''),
    crypt(@password::text, gen_salt('bf')),
    @gender::gender_type,
    nullif(@about::text, ''),
    nullif(@location::text, ''),
    nullif(@looking_for::text, '')
)
returning id;

-- name: UpdateLastLogin :exec
update users set last_login_at = (current_timestamp at time zone 'utc') where id = @id;

-- name: GetMyAccount :one
select
    coalesce(picture, 'none') as picture,
    username,
    email,
    coalesce(about, '') as about,
    coalesce(location, '') as location,
    coalesce(looking_for, '') as looking_for
from users
where id = @id;

-- name: UpdateMyAccount :one
update users
set
    username = case
        when sqlc.narg('username') is null or nullif(sqlc.narg('username')::text, '') is null then username
        else sqlc.narg('username')::text
    end,
    about = case
        when sqlc.narg('about') is null then about
        else nullif(sqlc.narg('about')::text, '')
    end,
    location = case
        when sqlc.narg('location') is null then location
        else nullif(sqlc.narg('location')::text, '')
    end,
    looking_for = case
        when sqlc.narg('looking_for') is null then looking_for
        else nullif(sqlc.narg('looking_for')::text, '')
    end
where id = @id
returning
    coalesce(picture, 'none') as picture,
    username,
    email,
    coalesce(about, '') as about,
    coalesce(location, '') as location,
    coalesce(looking_for, '') as looking_for;