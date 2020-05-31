
-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER,
    "name" TEXT,
    "email" TEXT
);

-- +migrate Down
DROP TABLE IF EXISTS "users";
