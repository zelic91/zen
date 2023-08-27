
CREATE TABLE devices (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL,
    platform TEXT,
    device_token TEXT,
    status TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    category_id BIGINT,
    status TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_categories_unique_title ON categories (title);

CREATE TABLE objects (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    category_id BIGINT,
    status TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
            REFERENCES categories(id)
            ON DELETE RESTRICT
);

-- migrate:down
DROP TABLE devices;
DROP TABLE objects;
DROP TABLE categories;
DROP TABLE users;