CREATE TABLE IF NOT EXISTS transaction_list (
    id SERIAL PRIMARY KEY ,
    amount INTEGER,
    qty INTEGER,
    total INTEGER,
    transaction_id INTEGER,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON UPDATE CASCADE ON DELETE SET NULL,
    product_id INTEGER,
    FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE SET NULL
)