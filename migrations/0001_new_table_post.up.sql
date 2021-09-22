CREATE TABLE post
(
    id      varchar(256),
    profile_id varchar(256),
    content text,
    FOREIGN KEY (profile_id) REFERENCES profile (uuid)
);