create extension if not exists "uuid-ossp";
create extension if not exists pgcrypto;

--
create table users (
    id uuid primary key default uuid_generate_v4(),
    picture varchar(255) default 'none',
    full_name varchar(255) not null,
    created_at timestamp default current_timestamp not null,
    -- auth fields
    email varchar(255) unique,
    google_uid varchar(255) unique,
    github_uid varchar(255) unique,
    phone_number varchar(15) unique,
    password varchar(255),
    -- github token
    github_username text default 'none'
    github_token text default 'none'
);

-- Create an index on email to improve search performance if necessary
create index idx_users_email on users(email);
create index idx_users_email on users(google_uid);
create index idx_auth_github_uid on users(github_uid);
create index idx_auth_phone_number on users(phone_number);

-- teams
create table wallet (
    id uuid primary key default uuid_generate_v4(),
    currency varchar(5) default 'KES',
    created_at timestamp default current_timestamp not null,

    -- 
    user_id UUID not null unique references users(id)  
);