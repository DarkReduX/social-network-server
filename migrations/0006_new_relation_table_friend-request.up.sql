CREATE TABLE friend_request
(
    from_user_id  varchar(256) REFERENCES profile (username) not null,
    to_user_id    varchar(256) REFERENCES profile (username) not null,
    request_state request_state                              not null,
    PRIMARY KEY (from_user_id, to_user_id)
)