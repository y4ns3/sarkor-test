-- +goose Up
    CREATE TABLE IF NOT EXISTS products(
        id serial PRIMARY KEY ,
        name VARCHAR(32) NOT NULL unique,
        description TEXT,
        price bigint NOT NULL,
        quantity INT NOT NULL
    );
-- +goose Down
DROP TABLE IF EXISTS products;