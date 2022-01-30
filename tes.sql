CREATE OR REPLACE PROCEDURE updateCoin()
LANGUAGE plpgsql
AS
$$
DECLARE 
   t_row users%rowtype;
BEGIN
    FOR t_row in SELECT * FROM users LOOP
        if t_row.id NOT IN (SELECT user_id FROM book_transaction WHERE user_id = t_row.id) then
        update users
            set coin = 10
        where id = t_row.id;
        end if;
    END LOOP;
END;
$$;