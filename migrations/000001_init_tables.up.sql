CREATE TABLE IF NOT EXISTS Users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    login VARCHAR(32) NOT NULL,
    passhash VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS Topic
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    text TEXT NOT NULL,
    header VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS Comments
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    topic_id INT NOT NULL,
    parent_id INT,
    text TEXT NOT NULL
);