CREATE TABLE profile
(
    username varchar(32),
    password varchar(256),
    avatar_link     text,
    last_activity   timestamp,
    created_at      timestamp,
    created_from_ip varchar(128),
    deleted_at      timestamp,
--  is_active -> false if deleted profile by user/true if profile is not deleted by user
    is_active       bool,
    PRIMARY KEY (username)
);