-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id            BIGSERIAL PRIMARY KEY,
    title         TEXT  NOT NULL,
    description   TEXT,
    user_id       INT   NOT NULL    REFERENCES users(id) ON DELETE CASCADE,
    start_at      TIMESTAMP WITH TIME ZONE  NOT NULL,
    duration      INTERVAL  NOT NULL,
    notify_offset INTERVAL,
    notified_at   TIMESTAMP WITH TIME ZONE,
    created_at    TIMESTAMP WITH TIME ZONE  NOT NULL    DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE  NOT NULL    DEFAULT NOW()
);

CREATE TRIGGER trg_set_updated_at BEFORE
UPDATE
ON events
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
