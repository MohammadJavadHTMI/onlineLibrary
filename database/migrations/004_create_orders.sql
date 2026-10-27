CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total_price INTEGER NOT NULL CHECK(total_price >= 0),status TEXT NOT NULL DEFAULT 'pending'CHECK(status IN ('pending', 'paid', 'cancelled', 'failed')),
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)REFERENCES users(id)ON DELETE RESTRICT
);