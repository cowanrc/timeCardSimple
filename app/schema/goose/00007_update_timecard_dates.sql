-- +goose Up
-- +goose StatementBegin
ALTER TABLE timecard
    ADD COLUMN duration NUMERIC DEFAULT 0,
    ADD COLUMN week_start_date DATE,
    ADD COLUMN bi_weekly_period_start DATE;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------

-- +goose Down
-- +goose StatementBegin
ALTER TABLE timecard 
    DROP COLUMN duration,
    DROP COLUMN week_start_date,
    DROP COLUMN bi_weekly_period_start;
-- +goose StatementEnd
