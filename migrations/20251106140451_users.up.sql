CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                        password TEXT NOT NULL,
                       points INT DEFAULT 0,
                       referrer_id INTEGER DEFAULT NULL REFERENCES users(id) ON DELETE SET NULL
);

-- INSERT INTO users(username, password)
-- VALUES ('Katya', 'katya'),('Dima', 'dima'), ('Irina', 'irina'), ('Misha', 'misha');