-- +goose Up 
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp"
CREATE TABLE victims (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    page VARCHAR(20) NOT NULL,
    user_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (username, page, password)
);
-- +goose Down
DROP TABLE victims;