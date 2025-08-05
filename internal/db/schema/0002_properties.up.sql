create type property_type as enum ('home', 'land', 'commercial');
create type property_status as enum ('available', 'sold', 'rented');
create type property_sale_type as enum ('for_sale', 'for_rent');

create table properties (
    id uuid primary key default uuid_generate_v4(),
    property_type property_type not null,
    status property_status default 'available',
    sale_type property_sale_type default 'for_sale',
    images text[] not null,
    title varchar(255) not null,
    description text,
    price decimal(12,2) not null,
    address text not null,
    city varchar(100) not null,
    state varchar(50) not null,
    bedrooms int default 0,
    bathrooms int default 0,
    square_feet int default 0,
    agent_id uuid not null references user_agents(id) on delete cascade,
    created_at timestamp not null default (current_timestamp at time zone 'utc')
);


create table subscriptions (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    tier varchar(20) not null check (tier in ('free', 'basic', 'pro')),
    status varchar(20) default 'active' check (status in ('active', 'cancelled', 'expired')),
    max_properties int not null,
    features text[], -- array of features included
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

-- Indexes for better performance
create index idx_properties_user_id on properties(user_id);
create index idx_properties_status on properties(status);
create index idx_properties_city on properties(city);
create index idx_properties_price on properties(price);
create index idx_subscriptions_user_id on subscriptions(user_id);
create index idx_subscriptions_tier on subscriptions(tier);
create index idx_subscriptions_status on subscriptions(status);

-- Insert default free subscription for existing users
insert into subscriptions (user_id, tier, max_properties, features)
select id, 'free', 1, array['Basic property listing', 'Standard support'] from users; 