DROP TABLE IF EXISTS students;

CREATE TABLE students (
    id varchar(32) PRIMARY KEY,
    name varchar(225) not null,
    age int not null,
);