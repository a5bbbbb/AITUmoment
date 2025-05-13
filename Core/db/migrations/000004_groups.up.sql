CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    educational_program_id SMALLINT NOT NULL,
    year SMALLINT NOT NULL,
    number SMALLINT NOT NULL,
    CONSTRAINT fk_educational_program 
        FOREIGN KEY (educational_program_id)
        REFERENCES educational_programs(id)
        ON DELETE RESTRICT,
    CONSTRAINT unique_group 
        UNIQUE (educational_program_id, year, number)
);
