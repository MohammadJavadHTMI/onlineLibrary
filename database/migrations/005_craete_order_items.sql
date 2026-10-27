CREATE TABLE IF NOT EXISTS order_items (
    id INTEGER PRIMARY KEY,
    order_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    price INTEGER NOT NULL CHECK(price >= 0),
    FOREIGN KEY (order_id)REFERENCES orders(id)ON DELETE CASCADE,
    FOREIGN KEY (book_id)REFERENCES books(id)ON DELETE RESTRICT,
    UNIQUE(order_id, book_id)
);