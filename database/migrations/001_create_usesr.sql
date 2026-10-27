CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('Admin', 'Customer', 'Author')),
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);