-- name: create_profile :one
insert into profiles (user_id, about, location, looking_for, birth_date, height, education, job_title)
values (@user_id::uuid, @about::text, @location::text, @looking_for::looking_for_type, @birth_date::date, @height::int, @education::text, @job_title::text)
returning id;

-- name: upsert_profile :one
insert into profiles (user_id, about, location, looking_for, birth_date, height, education, job_title)
values (@user_id::uuid, @about::text, @location::text, @looking_for::looking_for_type, @birth_date::date, @height::int, @education::text, @job_title::text)
on conflict (user_id)
do update set
    about = excluded.about,
    location = excluded.location,
    looking_for = excluded.looking_for,
    birth_date = excluded.birth_date,
    height = excluded.height,
    education = excluded.education,
    job_title = excluded.job_title
returning id;

-- name: add_profile_photo :one
insert into profile_photos (user_id, image_url, position, is_primary)
values (@user_id::uuid, @image_url::text, @position::int, @is_primary::boolean)
returning id;

-- name: like_user :one
insert into likes (from_user_id, to_user_id)
values (@from_user_id::uuid, @to_user_id::uuid)
on conflict (from_user_id, to_user_id)
do update set created_at = (current_timestamp at time zone 'utc')
returning id;

-- name: create_match :one
insert into matches (user_one_id, user_two_id)
values (@user_one_id::uuid, @user_two_id::uuid)
returning id;

-- name: create_conversation :one
insert into conversations (user_one_id, user_two_id)
values (@user_one_id::uuid, @user_two_id::uuid)
returning id;

-- name: create_message :one
insert into messages (conversation_id, sender_id, message_text)
values (@conversation_id::uuid, @sender_id::uuid, @message_text::text)
returning id;

-- name: list_conversation_messages :many
select id, conversation_id, sender_id, message_text, is_read, created_at
from messages
where conversation_id = @conversation_id::uuid
order by created_at asc;

-- name: create_story :one
insert into stories (user_id, media_url, media_type, expires_at)
values (@user_id::uuid, @media_url::text, @media_type::story_media_type, @expires_at::timestamp)
returning id;

-- name: block_user :one
insert into blocks (blocker_id, blocked_id)
values (@blocker_id::uuid, @blocked_id::uuid)
on conflict (blocker_id, blocked_id)
do update set created_at = (current_timestamp at time zone 'utc')
returning id;

-- name: report_user :one
insert into reports (reporter_id, reported_user_id, reason)
values (@reporter_id::uuid, @reported_user_id::uuid, @reason::text)
returning id;

-- name: ListFeed :many
select
    u.id as user_id,
    u.username,
    coalesce(u.picture, 'none') as picture,
    coalesce(p.about, '') as about,
    coalesce(p.location, '') as location,
    coalesce(p.looking_for::text, '') as looking_for,
    p.birth_date,
    coalesce(ph.image_url, '') as photo_url,
    case when sqlc.arg(can_view_phone)::boolean then coalesce(u.phone, '') else '' end as phone
from users u
join profiles p on p.user_id = u.id
left join lateral (
    select image_url
    from profile_photos
    where user_id = u.id
    order by is_primary desc, position asc
    limit 1
) ph on true
where u.id <> sqlc.arg(user_id)::uuid
  and u.status = 'active'
order by u.created_at desc
limit 15;

-- name: LikeAndMaybeMatch :one
with like_ins as (
    insert into likes (from_user_id, to_user_id)
    values (sqlc.arg(from_user_id)::uuid, sqlc.arg(to_user_id)::uuid)
    on conflict (from_user_id, to_user_id)
    do update set created_at = (current_timestamp at time zone 'utc')
    returning 1
),
mutual as (
    select exists(
        select 1
        from likes
        where from_user_id = sqlc.arg(to_user_id)::uuid
          and to_user_id = sqlc.arg(from_user_id)::uuid
    ) as matched
),
match_ins as (
    insert into matches (user_one_id, user_two_id)
    select least(sqlc.arg(from_user_id)::uuid, sqlc.arg(to_user_id)::uuid), greatest(sqlc.arg(from_user_id)::uuid, sqlc.arg(to_user_id)::uuid)
    where (select matched from mutual)
    on conflict do nothing
    returning 1
)
select (select matched from mutual) as matched;

-- name: ListActiveStories :many
select s.id, s.user_id, u.username, s.media_url, s.media_type, s.expires_at, s.created_at
from stories s
join users u on u.id = s.user_id
where s.expires_at > (current_timestamp at time zone 'utc')
  and s.user_id <> sqlc.arg(user_id)::uuid
order by s.created_at desc
limit 15;

-- name: UpsertDatingInteraction :one
insert into dating_interactions (user_id, target_user_id, action)
values (sqlc.arg(user_id)::uuid, sqlc.arg(target_user_id)::uuid, sqlc.arg(action)::dating_interaction_action)
on conflict (user_id, target_user_id)
do update set
    action = excluded.action,
    acted_at = (current_timestamp at time zone 'utc')
returning id, user_id, target_user_id, action, acted_at;

-- name: ListDatingHistory :many
select
    di.id,
    di.user_id,
    di.target_user_id,
    di.action,
    di.acted_at,
    u.username,
    coalesce(u.picture, 'none') as picture,
    coalesce(p.about, '') as about,
    coalesce(p.location, '') as location,
    coalesce(p.looking_for::text, '') as looking_for,
    coalesce(p.birth_date::text, '') as birth_date,
    coalesce(ph.image_url, '') as photo_url
from dating_interactions di
join users u on u.id = di.target_user_id
left join profiles p on p.user_id = u.id
left join lateral (
    select image_url
    from profile_photos
    where user_id = u.id
    order by is_primary desc, position asc
    limit 1
) ph on true
where di.user_id = sqlc.arg(user_id)::uuid
order by di.acted_at desc
limit 50;
