CREATE DATABASE IF NOT EXISTS  dev;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL,
  identifier VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tweets (
  id serial NOT NULL,
  identifier varchar(255) NOT NULL,
  text VARCHAR(255) NOT NULL,
  USER_id INT REFERENCES users(id),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

-- sample below
-- INSERT INTO users VALUES (1, 'dklajf', 'icchy', now());
