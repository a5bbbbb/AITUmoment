ALTER TABLE upvotes 
ADD CONSTRAINT unique_user_thread 
UNIQUE (userID, threadID);
