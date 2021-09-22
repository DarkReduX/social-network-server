CREATE OR REPLACE FUNCTION create_profile(in_uuid profile.uuid%TYPE, in_username profile.username%TYPE,
                                          in_password profile.password%TYPE,
                                          in_avatar_link profile.avatar_link%TYPE,
                                          in_last_activity profile.last_activity%TYPE,
                                          in_created_at profile.created_at%TYPE,
                                          in_created_from_ip profile.created_from_ip%TYPE)
    RETURNS VOID AS
$$
BEGIN
    INSERT INTO profile (uuid,
                         username,
                         password,
                         avatar_link,
                         last_activity,
                         created_at,
                         created_from_ip,
                         deleted_at,
                         is_active)
    VALUES (in_uuid, in_username, in_password, in_avatar_link, in_last_activity, in_created_at, in_created_from_ip,
            null, true);
END
$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION get_profile(in_uuid profile.uuid%TYPE) RETURNS SETOF profile AS
$$
BEGIN
    RETURN QUERY SELECT * FROM profile WHERE uuid = in_uuid;
END
$$
    LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION update_profile(in_uuid profile.uuid%TYPE, in_username profile.username%TYPE,
                                          in_password profile.password%TYPE,
                                          in_avatar_link profile.avatar_link%TYPE,
                                          in_last_activity profile.last_activity%TYPE,
                                          in_deleted_at profile.deleted_at%TYPE,
                                          in_is_active profile.is_active%TYPE) RETURNS VOID AS
$$
BEGIN
    UPDATE profile
    SET username      = in_username,
        password      = in_password,
        avatar_link   = in_avatar_link,
        last_activity = in_last_activity,
        deleted_at    = in_deleted_at,
        is_active     = in_is_active
    WHERE uuid = in_uuid;
END
$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION delete_profile(in_uuid profile.uuid%TYPE) RETURNS VOID AS
$$
BEGIN
    UPDATE profile SET is_active = false, deleted_at = now() WHERE uuid = in_uuid;
END
$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION get_profile_by_name(in_username profile.username%TYPE) RETURNS SETOF profile AS
$$
BEGIN
    RETURN QUERY SELECT * FROM profile WHERE username = in_username;
END
$$ LANGUAGE 'plpgsql';

