create extension if not exists "uuid-ossp";
create extension if not exists pgcrypto;

create type gender_type as enum ('male', 'female');
create type user_status as enum ('active', 'suspended', 'deleted');

create or replace function set_updated_at()
returns trigger as $$
begin
    new.updated_at = (current_timestamp at time zone 'utc');
    return new;
end;
$$ language plpgsql;

create table if not exists users (
    id uuid primary key default uuid_generate_v4(),
    username varchar(50) not null,
    email varchar(255) not null,
    phone varchar(15),
    password_hash varchar(255),
    gender gender_type,
    is_verified boolean not null default false,
    last_login_at timestamp,
    status user_status not null default 'active',
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc'),
    google_uid varchar(255) unique,
    picture varchar(255) default 'none',
    about text,
    location varchar(255),
    looking_for varchar(255),
    constraint users_username_unique unique (username),
    constraint users_email_unique unique (email),
    constraint users_phone_unique unique (phone)
);

drop trigger if exists trg_users_updated_at on users;
create trigger trg_users_updated_at
before update on users
for each row
execute function set_updated_at();

create index idx_users_status on users(status);
create index idx_users_last_login_at on users(last_login_at);

