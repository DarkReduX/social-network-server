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

CREATE TABLE post
(
    id      varchar(256),
    profile_id varchar(256),
    content text,
    FOREIGN KEY (profile_id) REFERENCES profile (username)
);

CREATE TABLE chat_room
(
    id varchar(256),
    PRIMARY KEY (id)
);

CREATE TABLE message
(
    id      varchar(256),
    user_id varchar(256),
    chat_id varchar(256),
    content text,
    PRIMARY KEY (id),
    FOREIGN KEY (chat_id) REFERENCES chat_room (id),
    FOREIGN KEY (user_id) REFERENCES profile (username)
);

CREATE TABLE chat_user
(
    chat_id varchar(256) REFERENCES chat_room (id),
    user_id varchar(256) REFERENCES profile (username),
    PRIMARY KEY (chat_id, user_id)
);


