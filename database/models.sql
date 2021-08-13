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

CREATE TABLE IF NOT EXISTS "user" (
    id serial NOT NULL,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    type_id int NOT NULL,
    career_id int NOT NULL,
    status Boolean not null,
    CONSTRAINT pk_user PRIMARY KEY(id),
    CONSTRAINT fk_user_career FOREIGN KEY(career_id) REFERENCES career(id),
    CONSTRAINT fk_user_type FOREIGN KEY(type_id) REFERENCES type(id)
    );

