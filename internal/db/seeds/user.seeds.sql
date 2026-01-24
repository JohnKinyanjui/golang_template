-- seed data for users table
-- all passwords are 'password123' when using email authentication

-- male users
insert into users (username, email, phone, password_hash, gender, about, location, looking_for, is_verified, picture)
values 
    ('johndoe', 'john.doe@example.com', '+1234567890', crypt('password123', gen_salt('bf')), 'male', 'Adventure seeker and outdoor enthusiast. Love hiking and exploring new places.', 'New York, USA', 'Serious relationship', true, 'https://picsum.photos/seed/johndoe/200/200.jpg'),
    
    ('mike_smith', 'mike.smith@example.com', '+1234567891', crypt('password123', gen_salt('bf')), 'male', 'Tech entrepreneur passionate about innovation and startups.', 'San Francisco, USA', 'Casual dating', true, 'https://picsum.photos/seed/mikesmith/200/200.jpg'),
    
    ('alex_wilson', 'alex.wilson@example.com', null, crypt('password123', gen_salt('bf')), 'male', 'Fitness fanatic and personal trainer. Helping others achieve their goals.', 'Los Angeles, USA', 'New friends', false, 'https://picsum.photos/seed/alexwilson/200/200.jpg'),
    
    ('david_brown', 'david.brown@example.com', '+1234567893', crypt('password123', gen_salt('bf')), 'male', 'Musician and producer. Jazz is my life!', 'Chicago, USA', 'Serious relationship', true, 'https://picsum.photos/seed/davidbrown/200/200.jpg'),
    
    ('chris_taylor', 'chris.taylor@example.com', null, crypt('password123', gen_salt('bf')), 'male', 'Chef exploring cuisines from around the world.', 'Miami, USA', 'Casual dating', false, 'https://picsum.photos/seed/christaylor/200/200.jpg'),

-- female users
    ('sarah_jones', 'sarah.jones@example.com', '+1234567894', crypt('password123', gen_salt('bf')), 'female', 'Yoga instructor and wellness coach. Finding balance in life.', 'Seattle, USA', 'Serious relationship', true, 'https://picsum.photos/seed/sarahjones/200/200.jpg'),
    
    ('emma_davis', 'emma.davis@example.com', '+1234567895', crypt('password123', gen_salt('bf')), 'female', 'Digital nomad and content creator. Currently exploring Southeast Asia.', 'Bangkok, Thailand', 'New friends', true, 'https://picsum.photos/seed/emmadavis/200/200.jpg'),
    
    ('lisa_miller', 'lisa.miller@example.com', null, crypt('password123', gen_salt('bf')), 'female', 'Art therapist and painter. Expressing emotions through colors.', 'Portland, USA', 'Casual dating', false, 'https://picsum.photos/seed/lisamiller/200/200.jpg'),
    
    ('jennifer_garcia', 'jennifer.garcia@example.com', '+1234567897', crypt('password123', gen_salt('bf')), 'female', 'Lawyer by day, salsa dancer by night.', 'Austin, USA', 'Serious relationship', true, 'https://picsum.photos/seed/jennifergarcia/200/200.jpg'),
    
    ('amanda_white', 'amanda.white@example.com', null, crypt('password123', gen_salt('bf')), 'female', 'Veterinarian with a love for all animals. Dog mom to three rescues.', 'Denver, USA', 'Casual dating', false, 'https://picsum.photos/seed/amandawhite/200/200.jpg'),

-- users with Google authentication (no password)
    ('google_user_1', 'google.user1@gmail.com', '+1234567899', null, 'male', 'Product manager at a tech giant. Always building the next big thing.', 'Mountain View, USA', 'Serious relationship', true, 'https://picsum.photos/seed/googleuser1/200/200.jpg'),
    
    ('google_user_2', 'google.user2@gmail.com', null, null, 'female', 'UX designer passionate about creating beautiful and functional interfaces.', 'San Francisco, USA', 'New friends', true, 'https://picsum.photos/seed/googleuser2/200/200.jpg');

