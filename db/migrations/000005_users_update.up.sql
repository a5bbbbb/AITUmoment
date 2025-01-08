-- First, add the new columns
ALTER TABLE users 
    ADD COLUMN email VARCHAR(255) NOT NULL,
    ADD COLUMN password VARCHAR(255) NOT NULL,
    ADD COLUMN public_name VARCHAR(100) NOT NULL,
    ADD COLUMN bio VARCHAR(100) NOT NULL,
    ADD COLUMN group_id INTEGER;

-- Add unique constraint for email
ALTER TABLE users 
    ADD CONSTRAINT users_email_unique UNIQUE (email);

-- Add foreign key constraint for group_id
ALTER TABLE users
    ADD CONSTRAINT fk_group_id 
        FOREIGN KEY (group_id) 
        REFERENCES groups(id)
        ON DELETE RESTRICT;
