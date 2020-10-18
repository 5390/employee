CREATE TABLE employee(
id serial PRIMARY KEY,
name VARCHAR(50) NOT NULL,
address VARCHAR(255) NOT NULL,
department VARCHAR(255),
skills VARCHAR(255) NOT NULL,
is_deleted int default 0
);