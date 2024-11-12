-- +goose Up
-- +goose StatementBegin
ALTER TABLE timecard
ALTER COLUMN start_time DROP NOT NULL,
ALTER COLUMN end_time DROP NOT NULL;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------


-- +goose Down
-- +goose StatementBegin
ALTER TABLE timecard
ALTER COLUMN start_time SET NOT NULL,
ALTER COLUMN end_time SET NOT NULL;
-- +goose StatementEnd
