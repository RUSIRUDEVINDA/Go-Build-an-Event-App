-- create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- unique number
    email TEXT NOT NULL UNIQUE,            -- email, duplicate not allowed
    name TEXT NOT NULL,                    -- name
    password TEXT NOT NULL                 -- password (hashed)
);