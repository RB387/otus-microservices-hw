-- migrate:up
CREATE TABLE users (
   username 	varchar(128) PRIMARY KEY,
   firstname 	varchar(128) NOT NULL,
   lastname	    varchar(128) NOT NULL,
   email		varchar(256) NOT NULL,
   phone		varchar(256) NOT NULL
);
-- migrate:down
DROP TABLE IF EXISTS users
