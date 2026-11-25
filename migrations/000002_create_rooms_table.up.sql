CREATE TABLE rooms (
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name         TEXT NOT NULL,                          
    age_limit    INT NOT NULL DEFAULT 0,
    description  TEXT NOT NULL,
    member_limit INT NOT NULL,
    is_dm        BOOLEAN NOT NULL DEFAULT false,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (NOT is_dm OR member_limit = 2)
);