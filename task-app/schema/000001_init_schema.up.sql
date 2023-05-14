CREATE TABLE  "task" (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)