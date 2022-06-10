CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR,
    password VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    created_at BIGINT
);

CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users (id),
    title VARCHAR,
    description VARCHAR,
    deadline int,
    reminder_period int,
    created_at BIGINT
);
