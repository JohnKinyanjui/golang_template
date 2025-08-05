-- name: GetUser :one
select 
    u.picture, u.full_name, u.email,
    u.role::text as role, u.phone_number, u.created_at, 
    json_build_object(
        'id', a.id,
        'picture', a.picture,
        'name', a.name,
        'description', a.description,
        'phone_number', a.phone_number,
        'email', a.email
    ) as agent
from users u
left join user_agents a on u.id = a.user_id 
where u.id = @user_id;

-- name: SignInWithGoogle :one
select id from users where email = @email::text and google_uid = @google_uid::text;

-- name: AuthenticateByEmail :one
select id from users where email = @email::text and password = crypt(@password::text, password) limit 1;

-- name: UpdateGoogleDetails :one
update users set email = @email::text , google_uid = @google_uid::text where id = @id returning id;

-- name: CreateUserAgent :one
insert into user_agents (user_id, picture, name, description, phone_number, email)
values (@user_id::uuid, @picture::text, @name::text, @description::text, @phone_number::text, @email::text)
returning id;