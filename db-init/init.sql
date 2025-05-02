CREATE TABLE student (id integer, name varchar(255), surname varchar(255), mark int);

CREATE OR REPLACE PROCEDURE delete_student_by_id(student_id INT)
LANGUAGE plpgsql
AS $$
BEGIN
  DELETE FROM student WHERE id = student_id;

  IF NOT FOUND THEN
    RAISE NOTICE 'Студент с id % не найден.', student_id;
  ELSE
    RAISE NOTICE 'Студент с id % удалён.', student_id;
  END IF;
END;
$$;
