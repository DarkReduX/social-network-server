create table profile
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

create table post
(
    id      varchar(256),
    profile_id varchar(256),
    content text,
    foreign key (profile_id) references profile (username)
);

create table chat_room
(
    id varchar(256),
    primary key (id)
);

create table message
(
    id      varchar(256),
    user_id varchar(256),
    chat_id varchar(256),
    content text,
    primary key (id),
    foreign key (chat_id) references chat_room (id),
    foreign key (user_id) references profile (username)
);

create table chat_user
(
    chat_id varchar(256) references chat_room (id),
    user_id varchar(256) references profile (username),
    PRIMARY KEY (chat_id, user_id)
);


--
-- truncate table post, profile, message, chat_user;
--
-- drop table message;
-- drop table chat_user;
-- drop table chat_room;
-- drop table post;
--
--drop function get_profile(varchar)
--drop table profile;

