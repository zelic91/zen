-- migrate:up
CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    bio text
);

-- migrate:down
DROP TABLE authors;