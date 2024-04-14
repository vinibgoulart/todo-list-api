
-- +migrate Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL UNIQUE,
    priority VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS tasks;
