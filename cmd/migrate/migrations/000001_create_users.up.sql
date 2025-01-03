CREATE EXTENSION IF NOT EXISTS citext;


CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY,
                                                       first_name VARCHAR(255) NOT NULL,
                                                                               last_name VARCHAR(255) UNIQUE NOT NULL,
                                                                                                             email CITEXT UNIQUE NOT NULL,
                                                                                                                                 username VARCHAR(255) UNIQUE NOT NULL,
                                                                                                                                                              password bytea NOT NULL,
                                                                                                                                                                             created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW());