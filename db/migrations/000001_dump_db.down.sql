-- 000001_initial_schema.down.sql
-- Note: We rarely use down migrations in production, but it's good practice to have them
-- Only drop if you really need to revert the migration during development
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS educational_programs;
