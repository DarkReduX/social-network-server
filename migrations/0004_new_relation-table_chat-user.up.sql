CREATE TABLE chat_user
(
    chat_id varchar(256) REFERENCES chat_room (id),
    user_id varchar(256) REFERENCES profile (uuid),
    PRIMARY KEY (chat_id, user_id)
);