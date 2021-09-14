CREATE OR REPLACE FUNCTION add_friend_request(in_from_user_id friend_request.from_user_id%TYPE,
                                              in_to_user_id friend_request.to_user_id%TYPE) RETURNS VOID AS
$$
BEGIN
    IF in_to_user_id = in_from_user_id THEN
        RAISE EXCEPTION SQLSTATE '90000' USING MESSAGE = 'Same value in fields: in_from_user_id and in_to_user_id';
    END IF;

    INSERT INTO friend_request (from_user_id, to_user_id, request_state)
    VALUES (in_from_user_id, in_to_user_id, 'IN_PROCESS');
END

$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION process_friend_request(in_from_user_id friend_request.from_user_id%TYPE,
                                                  in_to_user_id friend_request.to_user_id%TYPE,
                                                  in_request_type friend_request.request_state%TYPE)
    RETURNS VOID AS
$$
BEGIN
    UPDATE friend_request
    SET request_state = in_request_type
    WHERE from_user_id = in_from_user_id
      AND to_user_id = in_to_user_id;
END
$$
    LANGUAGE 'plpgsql'