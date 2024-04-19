CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS lesson (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   name text,
   unit_id int,
   text text,
   duration int,
   user_id int,
   description text,
   grade_id int,
   subject_id int,
   image_key text,
   concept_id int,
   skill_id int,
   created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   active bool
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
    BEFORE UPDATE ON lesson
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

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
      material_id text,
      lesson_id text,
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

INSERT INTO lesson  ("name", "unit_id", "text", "duration", "user_id", "description", "grade_id", "subject_id","concept_id","skill_id","active") VALUES
                    ('How to count',1,'Title: Exploring Geometric Shapes

Lesson Overview:
In this engaging math lesson, students will delve into the fascinating world of geometric shapes. Through a series of interactive activities and hands-on explorations, they will develop a deeper understanding of fundamental geometric concepts and their applications in real-world scenarios.

Objective:

To introduce students to various geometric shapes and their properties.
To foster critical thinking and problem-solving skills through exploratory learning.
To encourage collaboration and communication among peers.
Lesson Content:

Introduction to Geometric Shapes:

Definition of geometric shapes.
Classification of shapes based on sides and angles.
Examples of common geometric shapes such as triangles, quadrilaterals, circles, and polygons.
Properties of Geometric Shapes:

Identification of key properties including sides, angles, and diagonals.
Exploration of symmetry, congruence, and similarity.
Investigation of special properties of specific shapes (e.g., equilateral triangles, rectangles).
Exploring 2D and 3D Shapes:

Differentiation between two-dimensional (2D) and three-dimensional (3D) shapes.
Visualization of 3D shapes through models and diagrams.
Discussion on volume, surface area, and other attributes of 3D shapes.
Practical Applications:

Real-life examples showcasing the use of geometric shapes in architecture, art, and design.
Problem-solving tasks involving geometric concepts (e.g., calculating areas, constructing shapes).
Activities:

Shape Scavenger Hunt:

Students search the classroom for various objects that represent different geometric shapes. They classify each object and discuss their findings.
Geometry Art Project:

Students create artwork using geometric shapes as building blocks. They apply their knowledge of shapes, symmetry, and patterns to design visually appealing compositions.
Mathematical Puzzles and Challenges:

Students work in groups to solve mathematical puzzles and challenges related to geometric shapes. They apply problem-solving strategies and share their solutions with the class.
Assessment:

Formative assessment through class discussions, observations, and student presentations.
Summative assessment through quizzes, worksheets, and project evaluations.
Conclusion:
By the end of this lesson, students will have gained a solid foundation in geometric concepts and their practical applications. They will be equipped with essential skills to analyze, interpret, and manipulate geometric shapes, paving the way for further exploration in the field of mathematics.',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,true),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,true),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,true),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,true),
                    ('How to count',1,'Here we will learn how to count,find everything that is hidden, open the bounderies',60,1,'Here we will learn how to count by throwing a ball',5,1,1,1,true),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,2,2,true),
                    ('How to Science',2,'Here we will learn how to Science',45,1,'Here we will learn how to Science by throwing a ball',4,2,2,2,true),
                    ('How to Science',1,'Here we will learn how to Science',60,1,'Here we will learn how to Science by throwing a ball',4,2,2,2,true),
                    ('How to count',1,'Here we will learn how to count',30,1,'Here we will learn how to count by throwing a ball',5,1,2,2,true),
                    ('How to Arts',1,'Here we will learn how to Arts',60,1,'Here we will learn how to Arts by throwing a ball',3,6,2,2,true),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,2,2,true),
                    ('How to count',1,'Here we will learn how to count',90,1,'Here we will learn how to count by throwing a ball',5,1,2,2,true),
                    ('How to count',1,'Here we will learn how to count',60,1,'Here we will learn how to count by throwing a ball',5,1,2,2,true);



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




