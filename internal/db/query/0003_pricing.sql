-- name: get_plan_by_code :one
select id, code, name, price_cents, currency, interval, max_profile_views_per_day, can_view_phone, created_at, updated_at
from subscription_plans
where code = @code::text;

-- name: get_active_subscription_for_user :one
select s.id, s.user_id, s.plan_id, s.status, s.started_at, s.current_period_end, s.canceled_at, s.created_at, s.updated_at,
       p.code as plan_code, p.max_profile_views_per_day, p.can_view_phone
from user_subscriptions s
join subscription_plans p on p.id = s.plan_id
where s.user_id = @user_id::uuid
  and s.status = 'active'
limit 1;

-- name: create_user_subscription :one
insert into user_subscriptions (user_id, plan_id, status, current_period_end)
values (@user_id::uuid, @plan_id::uuid, @status::subscription_status, @current_period_end::timestamp)
returning id;

-- name: upsert_daily_usage :one
insert into user_feature_usage (user_id, feature_code, period_start, used_count)
values (@user_id::uuid, @feature_code::text, @period_start::date, @used_count::int)
on conflict (user_id, feature_code, period_start)
do update set
    used_count = user_feature_usage.used_count + excluded.used_count
returning id;

-- name: get_daily_usage :one
select id, user_id, feature_code, period_start, used_count
from user_feature_usage
where user_id = @user_id::uuid
  and feature_code = @feature_code::text
  and period_start = @period_start::date;

-- name: GetActiveSubscriptionForUser :one
select s.id, s.user_id, s.plan_id, s.status, s.started_at, s.current_period_end, s.canceled_at, s.created_at, s.updated_at,
       p.code as plan_code, p.max_profile_views_per_day, p.can_view_phone
from user_subscriptions s
join subscription_plans p on p.id = s.plan_id
where s.user_id = @user_id::uuid
  and s.status = 'active'
limit 1;

-- name: GetDailyUsage :one
select id, user_id, feature_code, period_start, used_count
from user_feature_usage
where user_id = @user_id::uuid
  and feature_code = @feature_code::text
  and period_start = @period_start::date;

-- name: UpsertDailyUsage :one
insert into user_feature_usage (user_id, feature_code, period_start, used_count)
values (@user_id::uuid, @feature_code::text, @period_start::date, @used_count::int)
on conflict (user_id, feature_code, period_start)
do update set
    used_count = user_feature_usage.used_count + excluded.used_count
returning id;
