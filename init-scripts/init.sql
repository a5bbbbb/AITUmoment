-- CREATE DATABASE aitu_moment;
--
-- \c aitu_moment


SELECT 'CREATE DATABASE aitu_moment'
WHERE NOT EXISTS (
    SELECT FROM pg_database WHERE datname = 'aitu_moment'
)\gexec

