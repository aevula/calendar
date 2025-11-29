-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    id          SERIAL  PRIMARY KEY,
    name        TEXT    NOT NULL,
    queue       TEXT    NOT NULL,
    payload     JSONB   NOT NULL,
    status      TEXT    NOT NULL,
    attempt     INT     NOT NULL    DEFAULT 0,
    start_at    TIMESTAMP WITH TIME ZONE    NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT NOW()
);

CREATE TRIGGER trg_set_updated_at BEFORE
UPDATE
ON tasks
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd