create type dating_interaction_action as enum ('like', 'nope');

create table if not exists dating_interactions (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    target_user_id uuid not null references users(id) on delete cascade,
    action dating_interaction_action not null,
    acted_at timestamp not null default (current_timestamp at time zone 'utc'),
    constraint dating_interactions_unique_pair unique (user_id, target_user_id),
    constraint dating_interactions_no_self check (user_id <> target_user_id)
);

create index if not exists idx_dating_interactions_user_id on dating_interactions(user_id);
create index if not exists idx_dating_interactions_target_user_id on dating_interactions(target_user_id);
create index if not exists idx_dating_interactions_acted_at on dating_interactions(acted_at);
