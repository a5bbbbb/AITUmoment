ALTER TABLE users 
    ADD COLUMN verified BOOLEAN;

UPDATE users SET verified = false;