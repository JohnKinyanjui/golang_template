-- seed data for dating module tables

-- profiles
insert into profiles (user_id, about, location, looking_for, birth_date, height, education, job_title)
values
    ((select id from users where email = 'john.doe@example.com'), 'adventure seeker and outdoor enthusiast.', 'new york, usa', 'serious', '1993-04-12', 180, 'bachelor degree', 'product manager'),
    ((select id from users where email = 'mike.smith@example.com'), 'tech entrepreneur passionate about innovation.', 'san francisco, usa', 'casual', '1990-09-22', 178, 'master degree', 'founder'),
    ((select id from users where email = 'alex.wilson@example.com'), 'fitness fanatic and personal trainer.', 'los angeles, usa', 'friendship', '1995-01-08', 182, 'diploma', 'personal trainer'),
    ((select id from users where email = 'david.brown@example.com'), 'musician and producer. jazz is my life.', 'chicago, usa', 'serious', '1991-07-17', 176, 'bachelor degree', 'music producer'),
    ((select id from users where email = 'chris.taylor@example.com'), 'chef exploring cuisines from around the world.', 'miami, usa', 'casual', '1994-11-03', 179, 'culinary school', 'chef'),
    ((select id from users where email = 'sarah.jones@example.com'), 'yoga instructor and wellness coach.', 'seattle, usa', 'serious', '1992-03-30', 168, 'bachelor degree', 'yoga instructor'),
    ((select id from users where email = 'emma.davis@example.com'), 'digital nomad and content creator.', 'bangkok, thailand', 'friendship', '1996-06-14', 165, 'bachelor degree', 'content creator'),
    ((select id from users where email = 'lisa.miller@example.com'), 'art therapist and painter.', 'portland, usa', 'casual', '1993-12-09', 164, 'master degree', 'art therapist'),
    ((select id from users where email = 'jennifer.garcia@example.com'), 'lawyer by day, salsa dancer by night.', 'austin, usa', 'serious', '1991-02-25', 167, 'juris doctor', 'lawyer'),
    ((select id from users where email = 'amanda.white@example.com'), 'veterinarian and dog mom to three rescues.', 'denver, usa', 'casual', '1994-05-19', 166, 'doctor of veterinary medicine', 'veterinarian'),
    ((select id from users where email = 'google.user1@gmail.com'), 'product manager at a tech giant.', 'mountain view, usa', 'serious', '1989-10-10', 181, 'bachelor degree', 'product manager'),
    ((select id from users where email = 'google.user2@gmail.com'), 'ux designer passionate about functional interfaces.', 'san francisco, usa', 'friendship', '1992-08-08', 163, 'bachelor degree', 'ux designer');

-- profile photos (2 per user, position 0 is primary)
insert into profile_photos (user_id, image_url, position, is_primary)
values
    ((select id from users where email = 'john.doe@example.com'), 'https://picsum.photos/seed/johndoe_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'john.doe@example.com'), 'https://picsum.photos/seed/johndoe_2/600/800.jpg', 1, false),
    ((select id from users where email = 'sarah.jones@example.com'), 'https://picsum.photos/seed/sarahjones_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'sarah.jones@example.com'), 'https://picsum.photos/seed/sarahjones_2/600/800.jpg', 1, false),
    ((select id from users where email = 'mike.smith@example.com'), 'https://picsum.photos/seed/mikesmith_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'mike.smith@example.com'), 'https://picsum.photos/seed/mikesmith_2/600/800.jpg', 1, false),
    ((select id from users where email = 'emma.davis@example.com'), 'https://picsum.photos/seed/emmadavis_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'emma.davis@example.com'), 'https://picsum.photos/seed/emmadavis_2/600/800.jpg', 1, false),
    ((select id from users where email = 'david.brown@example.com'), 'https://picsum.photos/seed/davidbrown_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'david.brown@example.com'), 'https://picsum.photos/seed/davidbrown_2/600/800.jpg', 1, false),
    ((select id from users where email = 'jennifer.garcia@example.com'), 'https://picsum.photos/seed/jennifergarcia_primary/600/800.jpg', 0, true),
    ((select id from users where email = 'jennifer.garcia@example.com'), 'https://picsum.photos/seed/jennifergarcia_2/600/800.jpg', 1, false);

-- likes
insert into likes (from_user_id, to_user_id, created_at)
values
    ((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'), (current_timestamp at time zone 'utc') - interval '3 days'),
    ((select id from users where email = 'sarah.jones@example.com'), (select id from users where email = 'john.doe@example.com'), (current_timestamp at time zone 'utc') - interval '3 days'),
    ((select id from users where email = 'mike.smith@example.com'), (select id from users where email = 'emma.davis@example.com'), (current_timestamp at time zone 'utc') - interval '2 days'),
    ((select id from users where email = 'emma.davis@example.com'), (select id from users where email = 'mike.smith@example.com'), (current_timestamp at time zone 'utc') - interval '2 days'),
    ((select id from users where email = 'david.brown@example.com'), (select id from users where email = 'lisa.miller@example.com'), (current_timestamp at time zone 'utc') - interval '1 days'),
    ((select id from users where email = 'alex.wilson@example.com'), (select id from users where email = 'amanda.white@example.com'), (current_timestamp at time zone 'utc') - interval '10 hours');

-- matches (based on mutual likes)
insert into matches (user_one_id, user_two_id, matched_at)
values
    ((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'), (current_timestamp at time zone 'utc') - interval '3 days'),
    ((select id from users where email = 'mike.smith@example.com'), (select id from users where email = 'emma.davis@example.com'), (current_timestamp at time zone 'utc') - interval '2 days');

-- conversations (for matched pairs)
insert into conversations (user_one_id, user_two_id, created_at)
values
    ((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'), (current_timestamp at time zone 'utc') - interval '3 days'),
    ((select id from users where email = 'mike.smith@example.com'), (select id from users where email = 'emma.davis@example.com'), (current_timestamp at time zone 'utc') - interval '2 days');

-- messages
insert into messages (conversation_id, sender_id, message_text, is_read, created_at)
values
    (
        (select id
         from conversations
         where least(user_one_id, user_two_id) = least((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'))
           and greatest(user_one_id, user_two_id) = greatest((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'))
        ),
        (select id from users where email = 'john.doe@example.com'),
        'hey sarah! loved your profile. want to grab coffee sometime?',
        true,
        (current_timestamp at time zone 'utc') - interval '3 days' + interval '5 minutes'
    ),
    (
        (select id
         from conversations
         where least(user_one_id, user_two_id) = least((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'))
           and greatest(user_one_id, user_two_id) = greatest((select id from users where email = 'john.doe@example.com'), (select id from users where email = 'sarah.jones@example.com'))
        ),
        (select id from users where email = 'sarah.jones@example.com'),
        'hi john! coffee sounds great. how about this weekend?',
        true,
        (current_timestamp at time zone 'utc') - interval '3 days' + interval '20 minutes'
    ),
    (
        (select id
         from conversations
         where least(user_one_id, user_two_id) = least((select id from users where email = 'mike.smith@example.com'), (select id from users where email = 'emma.davis@example.com'))
           and greatest(user_one_id, user_two_id) = greatest((select id from users where email = 'mike.smith@example.com'), (select id from users where email = 'emma.davis@example.com'))
        ),
        (select id from users where email = 'emma.davis@example.com'),
        'hey mike! i saw you are into startups too. what are you building?',
        false,
        (current_timestamp at time zone 'utc') - interval '2 days' + interval '10 minutes'
    );

-- stories (expires 12 hours from now)
insert into stories (user_id, media_url, media_type, expires_at, created_at)
values
    ((select id from users where email = 'emma.davis@example.com'), 'https://picsum.photos/seed/story_emma_1/720/1280.jpg', 'image', (current_timestamp at time zone 'utc') + interval '12 hours', (current_timestamp at time zone 'utc') - interval '30 minutes'),
    ((select id from users where email = 'john.doe@example.com'), 'https://picsum.photos/seed/story_john_1/720/1280.jpg', 'image', (current_timestamp at time zone 'utc') + interval '12 hours', (current_timestamp at time zone 'utc') - interval '45 minutes');

-- blocks
insert into blocks (blocker_id, blocked_id, created_at)
values
    ((select id from users where email = 'chris.taylor@example.com'), (select id from users where email = 'alex.wilson@example.com'), (current_timestamp at time zone 'utc') - interval '6 hours');

-- reports
insert into reports (reporter_id, reported_user_id, reason, created_at)
values
    ((select id from users where email = 'lisa.miller@example.com'), (select id from users where email = 'david.brown@example.com'), 'spam messages', (current_timestamp at time zone 'utc') - interval '5 hours');