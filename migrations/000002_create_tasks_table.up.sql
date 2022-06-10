CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users (id),
    title VARCHAR,
    description VARCHAR,
    start_date BIGINT,
    end_date BIGINT,
    reminder_period BIGINT,
    created_at BIGINT
);
