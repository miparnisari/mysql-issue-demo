-- +goose Up
CREATE TABLE store (
    id CHAR(26) PRIMARY KEY,
    name VARCHAR(64) NOT NULL
);

-- +goose Down
DROP TABLE store;