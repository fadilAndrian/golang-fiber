CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY ,
    name VARCHAR,
    email VARCHAR,
    password VARCHAR
)