CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY,
    order_id INTEGER NOT NULL,
    amount INTEGER NOT NULL CHECK(amount >= 0),
    authority TEXT UNIQUE,
    ref_id TEXT UNIQUE, status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'success', 'failed')),
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    paid_at TEXT,
    FOREIGN KEY (order_id)REFERENCES orders(id)ON DELETE RESTRICT
);