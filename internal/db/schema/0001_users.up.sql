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
    picture varchar(255) default 'none',
    full_name varchar(50) not null,
    email varchar(255) unique not null,
    phone varchar(15) unique,
    password_hash varchar(255),
    status user_status not null default 'active',
    google_uid varchar(255) unique,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc')
);

drop trigger if exists trg_users_updated_at on users;
create trigger trg_users_updated_at
before update on users
for each row
execute function set_updated_at();

create index idx_users_status on users(status);
create index idx_users_last_login_at on users(last_login_at);
