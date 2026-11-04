CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role_id BIGINT NOT NULL DEFAULT 1 REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users (id, name, email, password_hash, role_id) VALUES
    (gen_random_uuid(), 'User', 'user1@example.com', crypt('password123', gen_salt('bf')), 1),
    (gen_random_uuid(), 'Mentor', 'user2@example.com', crypt('password123', gen_salt('bf')), 2),
    (gen_random_uuid(), 'Admin', 'user3@example.com', crypt('password123', gen_salt('bf')), 3)
ON CONFLICT (email) DO NOTHING;
