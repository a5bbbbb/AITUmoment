-- 000001_initial_schema.up.sql
CREATE TABLE IF NOT EXISTS educational_programs (
    id SMALLINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    educational_program SMALLINT NOT NULL DEFAULT 0,
    CONSTRAINT fk_educational_program 
        FOREIGN KEY (educational_program)
        REFERENCES educational_programs(id)
        ON DELETE RESTRICT
);
