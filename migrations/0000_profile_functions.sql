CREATE or replace function create_profile(varchar(32), varchar(256), varchar(256), timestamp, timestamp, varchar(128)) returns void as
$$
begin
    insert into profile (username,
                         password,
                         avatar_link,
                         last_activity,
                         created_at,
                         created_from_ip,
                         deleted_at,
                         is_active)
    values ($1, $2, $3, $4, $5, $6, null, true);
end
$$ language 'plpgsql';

CREATE or replace function get_profile(varchar(32)) returns setof profile as
$$
begin
    return query select * from profile where username = $1;
end
$$
    language 'plpgsql';

CREATE or replace function update_profile(varchar(32), varchar(256), varchar(256), timestamp,
                                          timestamp, bool) returns void as
$$
begin
    update profile
    set password        = $2,
        avatar_link     = $3,
        last_activity   = $4,
        deleted_at      = $5,
        is_active       = $6
    where username = $1;
end
$$ language 'plpgsql';

CREATE or replace function delete_profile(varchar(32)) returns void as
$$
begin
    update profile set is_active = false, deleted_at = now() where username = $1;
end
$$ language 'plpgsql';



