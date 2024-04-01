CREATE TABLE IF NOT EXISTS lesson (
   id SERIAL PRIMARY KEY,
   name text,
   user_id int,
   description text,
   grade_id int,
   subject_id int,
   created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user" (
    id integer PRIMARY KEY,
    name text
);

CREATE TABLE IF NOT EXISTS material (
    id integer PRIMARY KEY,
    material_type_id integer,
    link text
);


CREATE TABLE IF NOT EXISTS material_type (
     id integer PRIMARY KEY,
     name text
);


CREATE TABLE IF NOT EXISTS material_to_lesson (
      material_id int,
      lesson_id int,
      primary key (material_id,lesson_id)
);


CREATE TABLE IF NOT EXISTS grade (
  id integer PRIMARY KEY,
  name text
);


CREATE TABLE IF NOT EXISTS subject (
    id integer PRIMARY KEY,
    name text
);


INSERT INTO subject VALUES (1,'mathematics'), (2,'english');

