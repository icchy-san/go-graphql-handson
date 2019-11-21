#!/bin/bash

psql -U root -d dev << "EOSQL"
CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL,
  identifier VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
EOSQL

psql -U root -d dev << "EOSQL"
CREATE TABLE IF NOT EXISTS tweets (
  id serial NOT NULL,
  text VARCHAR(255) NOT NULL,
  USER_id INT REFERENCES users(id),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
EOSQL
