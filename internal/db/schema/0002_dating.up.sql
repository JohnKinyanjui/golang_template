create type looking_for_type as enum ('serious', 'casual', 'friendship', 'other');
create type story_media_type as enum ('image', 'video');

create table if not exists profiles (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null unique references users(id) on delete cascade,
    about text,
    location varchar(255),
    looking_for looking_for_type,
    birth_date date,
    height int,
    education varchar(255),
    job_title varchar(255),
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    updated_at timestamp not null default (current_timestamp at time zone 'utc')
);

drop trigger if exists trg_profiles_updated_at on profiles;
create trigger trg_profiles_updated_at
before update on profiles
for each row
execute function set_updated_at();

create index if not exists idx_profiles_user_id on profiles(user_id);
create index if not exists idx_profiles_location on profiles(location);
create index if not exists idx_profiles_looking_for on profiles(looking_for);

create table if not exists profile_photos (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    image_url text not null,
    position int not null default 0,
    is_primary boolean not null default false,
    created_at timestamp not null default (current_timestamp at time zone 'utc')
);

create index if not exists idx_profile_photos_user_id on profile_photos(user_id);
create unique index if not exists idx_profile_photos_user_position on profile_photos(user_id, position);
create unique index if not exists idx_profile_photos_one_primary on profile_photos(user_id) where is_primary;

create table if not exists likes (
    id uuid primary key default uuid_generate_v4(),
    from_user_id uuid not null references users(id) on delete cascade,
    to_user_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint likes_unique_pair unique (from_user_id, to_user_id),
    constraint likes_no_self check (from_user_id <> to_user_id)
);

create index if not exists idx_likes_from_user_id on likes(from_user_id);
create index if not exists idx_likes_to_user_id on likes(to_user_id);
create index if not exists idx_likes_created_at on likes(created_at);

create table if not exists matches (
    id uuid primary key default uuid_generate_v4(),
    user_one_id uuid not null references users(id) on delete cascade,
    user_two_id uuid not null references users(id) on delete cascade,
    matched_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint matches_no_self check (user_one_id <> user_two_id)
);

create unique index if not exists idx_matches_unique_pair
on matches (least(user_one_id, user_two_id), greatest(user_one_id, user_two_id));
create index if not exists idx_matches_matched_at on matches(matched_at);

create table if not exists conversations (
    id uuid primary key default uuid_generate_v4(),
    user_one_id uuid not null references users(id) on delete cascade,
    user_two_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint conversations_no_self check (user_one_id <> user_two_id)
);

create unique index if not exists idx_conversations_unique_pair
on conversations (least(user_one_id, user_two_id), greatest(user_one_id, user_two_id));
create index if not exists idx_conversations_created_at on conversations(created_at);

create table if not exists messages (
    id uuid primary key default uuid_generate_v4(),
    conversation_id uuid not null references conversations(id) on delete cascade,
    sender_id uuid not null references users(id) on delete cascade,
    message_text text not null,
    is_read boolean not null default false,
    created_at timestamp not null default (current_timestamp at time zone 'utc')
);

create index if not exists idx_messages_conversation_id on messages(conversation_id);
create index if not exists idx_messages_sender_id on messages(sender_id);
create index if not exists idx_messages_created_at on messages(created_at);
create index if not exists idx_messages_conversation_is_read on messages(conversation_id, is_read);

create table if not exists stories (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    media_url text not null,
    media_type story_media_type not null,
    expires_at timestamp not null,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint stories_expires_after_created check (expires_at > created_at)
);

create index if not exists idx_stories_user_id on stories(user_id);
create index if not exists idx_stories_expires_at on stories(expires_at);

create table if not exists blocks (
    id uuid primary key default uuid_generate_v4(),
    blocker_id uuid not null references users(id) on delete cascade,
    blocked_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint blocks_unique_pair unique (blocker_id, blocked_id),
    constraint blocks_no_self check (blocker_id <> blocked_id)
);

create index if not exists idx_blocks_blocker_id on blocks(blocker_id);
create index if not exists idx_blocks_blocked_id on blocks(blocked_id);

create table if not exists reports (
    id uuid primary key default uuid_generate_v4(),
    reporter_id uuid not null references users(id) on delete cascade,
    reported_user_id uuid not null references users(id) on delete cascade,
    reason text not null,
    created_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint reports_no_self check (reporter_id <> reported_user_id)
);

create index if not exists idx_reports_reporter_id on reports(reporter_id);
create index if not exists idx_reports_reported_user_id on reports(reported_user_id);
create index if not exists idx_reports_created_at on reports(created_at);
