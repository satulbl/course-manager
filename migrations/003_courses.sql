CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    price BIGINT NOT NULL CHECK (price >= 0),
    description TEXT NOT NULL,
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    mentor_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL CHECK (end_date > start_date),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);