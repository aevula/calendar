-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE events
    ADD COLUMN notify_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW();

UPDATE events
SET notify_at = start_at - COALESCE(notify_offset, INTERVAL '0');

ALTER TABLE events
    ALTER COLUMN notify_at DROP DEFAULT;

ALTER TABLE events
    DROP COLUMN notify_offset;

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;

ALTER TABLE events
    ADD COLUMN notify_offset INTERVAL;

UPDATE events
SET notify_offset = start_at - notify_at;

ALTER TABLE events
    DROP COLUMN notify_at;

COMMIT;
-- +goose StatementEnd
