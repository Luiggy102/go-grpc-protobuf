DROP TABLE IF EXISTS students;

CREATE TABLE students (
    id varchar(32) PRIMARY KEY,
    name varchar(225) not null,
    age int not null
);

DROP TABLE IF EXISTS tests;
CREATE TABLE tests (
    id varchar(32) PRIMARY KEY,
    name varchar(225) not null
);

DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
    id varchar(32) PRIMARY KEY,
    question varchar(225) not null,
    answer varchar(225) not null,
    test_id varchar(225) not null,
    FOREIGN KEY(test_id) REFERENCES tests(id)
);
