CREATE TABLE IF NOT EXISTS upvotes (
    id SERIAL PRIMARY KEY,
    userID SMALLINT NOT NULL,
    threadID SMALLINT NOT NULL,
    CONSTRAINT fk_user_id 
        FOREIGN KEY (userID)
        REFERENCES users(id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_thread_id 
        FOREIGN KEY (threadID)
        REFERENCES threads(thread_id)
        ON DELETE RESTRICT
);
