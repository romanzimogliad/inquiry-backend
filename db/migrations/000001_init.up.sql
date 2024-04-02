CREATE TABLE IF NOT EXISTS lesson (
   id SERIAL PRIMARY KEY,
   name text,
   unit_id int,
   text text,
   duration int,
   user_id int,
   description text,
   grade_id int,
   subject_id int,
   image_id int,
   concept_id int,
   skill_id int,
   created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS concept (
    id SERIAL PRIMARY KEY,
    name text
);


CREATE TABLE IF NOT EXISTS unit (
    id SERIAL PRIMARY KEY,
    name text
);

CREATE TABLE IF NOT EXISTS skill (
                                    id SERIAL PRIMARY KEY,
                                    name text
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

INSERT INTO lesson  ("name", "unit_id", "text", "duration", "user_id", "description", "grade_id", "subject_id", "image_id","concept_id","skill_id") VALUES
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,1),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,1),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,1),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,1),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,1),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,1,2,2),
                    ('How to Science',2,'Here we will learn how to Science',45,1,'Here we will learn how to Science by throwing a ball',4,2,1,2,2),
                    ('How to Science',1,'Here we will learn how to Science',60,1,'Here we will learn how to Science by throwing a ball',4,2,1,2,2),
                    ('How to count',1,'Here we will learn how to count',30,1,'Here we will learn how to count by throwing a ball',5,1,1,2,2),
                    ('How to Arts',1,'Here we will learn how to Arts',60,1,'Here we will learn how to Arts by throwing a ball',3,6,1,2,2),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,1,2,2),
                    ('How to count',1,'Here we will learn how to count',90,1,'Here we will learn how to count by throwing a ball',5,1,1,2,2),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,1,2,2);



INSERT INTO subject VALUES (0,'All'),(1,'Mathematics'), (2,'Science'), (3,'Humanities'),
                           (4,'Foreign Languages'),(5,'Physical Education'),
                           (6,'Arts'),(7,'Information Technology'),(8,'Design and Technology');

INSERT INTO concept VALUES (0,'All'),(1,'Form'), (2,'Function'), (3,'Causation'),
                           (4,'Change'),(5, 'Connection'),(6,'Perspective'),
                           (7, 'Responsibility'),(8,'Reflection');


INSERT INTO unit VALUES (0,'All'),(1,'Who we are'),(2,'Where we are in place and time'),
                        (3,'How we express ourselves'),
                        (4,'How the world works'),
                        (5,'How we organize ourselves'),
                        (6,'Sharing the planet');
INSERT INTO skill VALUES (0,'All'),(1,'Thinking'), (2,'Communication'), (3,'Research'),
                           (4,'Self-management'),(5, 'Social');




