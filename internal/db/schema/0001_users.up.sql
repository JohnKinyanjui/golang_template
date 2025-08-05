create extension if not exists "uuid-ossp";
create extension if not exists pgcrypto;

create type user_role as enum ('admin', 'user', 'guest');
create table if not exists users (
    id uuid primary key default uuid_generate_v4(),
    picture varchar(255) default 'none',
    role user_role default 'user',
    full_name varchar(255) not null,
    email varchar(255) unique,
    google_uid varchar(255) unique,
    phone_number varchar(15) unique,
    password varchar(255),
    created_at timestamp not null default (current_timestamp at time zone 'utc')
);

create table if not exists user_agents (
    id uuid primary key default uuid_generate_v4(),
    picture varchar(255) default 'none',
    name varchar(255) not null,
    description text not null,
    phone_number varchar(15) not null,
    email varchar(255) not null,
    user_id uuid unique references users(id) on delete cascade,
    created_at timestamp not null default (current_timestamp at time zone 'utc')
);

-- Create an index on email to improve search performance if necessary
create index idx_users_email on users(email);
create index idx_users_email on users(google_uid);
create index idx_auth_phone_number on users(phone_number);

