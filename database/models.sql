CREATE TABLE IF NOT EXISTS career (
    id serial NOT NULL,
    name VARCHAR(150) NOT NULL,
    status Boolean not null,
    CONSTRAINT pk_career PRIMARY KEY(id)
    );

CREATE TABLE IF NOT EXISTS type (
    id serial NOT NULL,
    name VARCHAR(150) NOT NULL,
    status Boolean not null,
    CONSTRAINT pk_type PRIMARY KEY(id)
    );

CREATE TABLE IF NOT EXISTS course (
    id serial NOT NULL,
    career_id int NOT NULL,
    name VARCHAR(150) NOT NULL,
    credits int NOT NULL,
    status Boolean not null,
    CONSTRAINT pk_course PRIMARY KEY(id),
    CONSTRAINT fk_course_career FOREIGN KEY(career_id) REFERENCES career(id)
);

