INSERT INTO educational_programs (id, name)
VALUES (0, 'NOT_SELECTED')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (1, 'Software Engineering')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (2, 'Computer Science')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (3, 'Big Data Analysis')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (4, 'Media Technologies')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (5, 'Mathematical and Computational Science')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (6, 'Big Data in Healthcare')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (7, 'Cybersecurity')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (8, 'Smart Technologies')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (9, 'Industrial Internet of Things')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (10, 'Electronic Engineering')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (11, 'IT Management')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (12, 'IT Entrepreneurship')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (13, 'AI Business')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

INSERT INTO educational_programs (id, name)
VALUES (14, 'Digital Journalism')
ON CONFLICT (id)
DO UPDATE SET name = EXCLUDED.name;

