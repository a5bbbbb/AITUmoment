-- Insert parent threads
INSERT INTO threads (creator_id, content, create_date, up_votes, parent_thread_id) VALUES
(7, 'Starting my journey in web development! Currently learning HTML and CSS basics. Any tips for beginners?', '2024-01-01 10:00:00+00', 15, NULL),
(7, 'Just completed my first responsive design project. The mobile-first approach really makes sense now!', '2024-01-02 11:30:00+00', 23, NULL),
(7, 'Question about JavaScript promises: What''s the best way to handle multiple async operations?', '2024-01-03 09:15:00+00', 31, NULL),
(7, 'Looking for study partners in our Algorithm Analysis course. Anyone interested?', '2024-01-04 14:20:00+00', 18, NULL),
(7, 'Share your favorite VS Code extensions! Mine is definitely GitLens.', '2024-01-05 16:45:00+00', 42, NULL),
(7, 'Just discovered the power of CSS Grid. Game changer for layouts!', '2024-01-06 08:30:00+00', 27, NULL),
(7, 'Working on my portfolio website. Should I use React or keep it simple with vanilla JS?', '2024-01-06 13:20:00+00', 35, NULL),
(7, 'Database design question: How do you handle many-to-many relationships effectively?', '2024-01-07 10:15:00+00', 29, NULL),
(7, 'Found a great resource for learning system design. Check it out: system-design-primer', '2024-01-07 15:40:00+00', 45, NULL),
(7, 'Successfully deployed my first Docker container! The learning curve was worth it.', '2024-01-08 09:10:00+00', 33, NULL);

-- Insert replies to various threads
INSERT INTO threads (creator_id, content, create_date, up_votes, parent_thread_id) VALUES
(7, 'Update: Started with CSS flexbox today. It''s amazing for handling one-dimensional layouts!', '2024-01-01 14:30:00+00', 12, 1),
(7, 'Found this great tutorial series on YouTube about responsive design patterns.', '2024-01-02 16:20:00+00', 8, 2),
(7, 'After some research, I found that async/await makes promise handling much cleaner.', '2024-01-03 11:45:00+00', 19, 3),
(7, 'I created a study schedule for algorithms if anyone wants to join!', '2024-01-04 17:30:00+00', 14, 4),
(7, 'Another great extension: Live Server. Makes local development so much easier.', '2024-01-05 18:20:00+00', 21, 5),
(7, 'Here''s a CodePen I made showcasing some cool CSS Grid layouts.', '2024-01-06 10:45:00+00', 16, 6),
(7, 'Decided to go with React. The component-based architecture is really growing on me.', '2024-01-06 15:50:00+00', 23, 7),
(7, 'Junction tables have been really helpful for handling many-to-many relationships.', '2024-01-07 12:30:00+00', 17, 8),
(7, 'Added some practical examples to the system design resources.', '2024-01-07 17:45:00+00', 25, 9),
(7, 'Wrote a simple guide for Docker beginners based on my learning experience.', '2024-01-08 11:20:00+00', 28, 10),
(7, 'Started learning about CSS animations. The transform property is fascinating!', '2024-01-01 16:45:00+00', 9, 1),
(7, 'Media queries are tricky but essential for responsive design.', '2024-01-02 18:30:00+00', 11, 2),
(7, 'Created a small demo project showing promise chaining patterns.', '2024-01-03 13:40:00+00', 15, 3),
(7, 'Found some great practice problems on LeetCode for our study group.', '2024-01-04 19:15:00+00', 13, 4),
(7, 'The Prettier extension has saved me so much time formatting code.', '2024-01-05 20:10:00+00', 18, 5),
(7, 'Exploring CSS Grid areas now - they make layout naming so intuitive!', '2024-01-06 12:25:00+00', 14, 6),
(7, 'Started learning about React hooks. useState and useEffect are amazing!', '2024-01-06 17:40:00+00', 20, 7),
(7, 'Implementing proper indexing really improved our query performance.', '2024-01-07 14:15:00+00', 16, 8),
(7, 'Added a section about microservices architecture to the resources.', '2024-01-07 19:30:00+00', 22, 9),
(7, 'Exploring Docker Compose for multi-container applications now.', '2024-01-08 13:40:00+00', 19, 10);
