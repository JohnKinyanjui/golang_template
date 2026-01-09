create type subscription_interval as enum ('monthly', 'yearly');
create type subscription_status as enum ('active', 'canceled', 'expired');

create table if not exists subscription_plans (
    id uuid primary key default uuid_generate_v4(),
    code varchar(50) not null unique,
    name varchar(100) not null,
    price_cents int not null default 0,
    currency varchar(10) not null default 'usd',
    interval subscription_interval not null default 'monthly',
    max_profile_views_per_day int,
    can_view_phone boolean not null default false,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc')
);

drop trigger if exists trg_subscription_plans_updated_at on subscription_plans;
create trigger trg_subscription_plans_updated_at
before update on subscription_plans
for each row
execute function set_updated_at();

create table if not exists user_subscriptions (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    plan_id uuid not null references subscription_plans(id) on delete restrict,
    status subscription_status not null default 'active',
    started_at timestamp not null default (current_timestamp at time zone 'utc'),
    current_period_end timestamp,
    canceled_at timestamp,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc')
);

drop trigger if exists trg_user_subscriptions_updated_at on user_subscriptions;
create trigger trg_user_subscriptions_updated_at
before update on user_subscriptions
for each row
execute function set_updated_at();

create index if not exists idx_user_subscriptions_user_id on user_subscriptions(user_id);
create index if not exists idx_user_subscriptions_status on user_subscriptions(status);
create unique index if not exists idx_user_subscriptions_one_active
on user_subscriptions(user_id)
where status = 'active';

create table if not exists user_feature_usage (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    feature_code varchar(50) not null,
    period_start date not null,
    used_count int not null default 0,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint user_feature_usage_unique unique (user_id, feature_code, period_start)
);

drop trigger if exists trg_user_feature_usage_updated_at on user_feature_usage;
create trigger trg_user_feature_usage_updated_at
before update on user_feature_usage
for each row
execute function set_updated_at();

create index if not exists idx_user_feature_usage_user_id on user_feature_usage(user_id);
create index if not exists idx_user_feature_usage_feature_period on user_feature_usage(feature_code, period_start);
