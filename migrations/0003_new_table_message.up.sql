CREATE TABLE message
(
    id      varchar(256),
    user_id varchar(256),
    chat_id varchar(256),
    content text,
    PRIMARY KEY (id),
    FOREIGN KEY (chat_id) REFERENCES chat_room (id),
    FOREIGN KEY (user_id) REFERENCES profile (uuid)
);