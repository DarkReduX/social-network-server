CREATE TABLE friend_request
(
    from_user_id  varchar(256) REFERENCES profile (uuid) not null,
    to_user_id    varchar(256) REFERENCES profile (uuid) not null,
    request_state request_state                              not null,
    PRIMARY KEY (from_user_id, to_user_id)
)