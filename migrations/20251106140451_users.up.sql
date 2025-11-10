CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(30) NOT NULL,
                        hashed_password CHAR(60) NOT NULL,
                       points INT DEFAULT 0,
                       referrer_id INTEGER DEFAULT NULL REFERENCES users(id) ON DELETE SET NULL
);
