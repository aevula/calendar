-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL  PRIMARY KEY,
    login       TEXT    NOT NULL    UNIQUE,
    name        TEXT    NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT NOW()
);

CREATE TRIGGER trg_set_updated_at BEFORE
UPDATE
ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd