CREATE TABLE IF NOT EXISTS threads (
    thread_id SERIAL PRIMARY KEY,
    creator_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    create_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    up_votes INTEGER DEFAULT 0,
    parent_thread_id INTEGER,
    FOREIGN KEY (creator_id) REFERENCES users(id),
    FOREIGN KEY (parent_thread_id) REFERENCES threads(thread_id),
    CHECK (up_votes >= 0)
);
