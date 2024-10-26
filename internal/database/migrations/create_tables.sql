CREATE TABLE IF NOT EXISTS categories (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT
);

CREATE TABLE IF NOT EXISTS customers (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT
);

CREATE TABLE IF NOT EXISTS  orders(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        customer_id INTEGER,
        date TEXT,
        FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS order_item (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        order_id INTEGER,
        item_id INTEGER,
        quantity INTEGER,
        FOREIGN KEY (order_id) REFERENCES orders(id),
        FOREIGN KEY (item_id) REFERENCES grocery_items(id)
);

CREATE TABLE IF NOT EXISTS grocery_items (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        unit_price REAL,
        stock INTEGER,
        category_id INTEGER,
        FOREIGN KEY (category_id) REFERENCES categories(id)
);
