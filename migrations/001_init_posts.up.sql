CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    allow_comments BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS posts_created_at_idx ON posts (created_at DESC);