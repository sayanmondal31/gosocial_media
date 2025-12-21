CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
     id bigserial PRIMARY KEY,
     username VARCHAR(255) UNIQUE NOT NULL,
     email citext UNIQUE NOT NULL, -- CITEXT case insensitive
     password bytea NOT NULL, -- bytea passwords store in hash format 
     create_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);