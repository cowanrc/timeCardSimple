-- +goose Up
-- +goose StatementBegin
ALTER TABLE weekly_summary
    ALTER COLUMN week_start_date DROP NOT NULL;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE pay_periods
    ALTER COLUMN period_start_date DROP NOT NULL,
    ALTER COLUMN period_end_date DROP NOT NULL;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------

-- +goose Down
-- +goose StatementBegin
ALTER TABLE weekly_summary
    ALTER COLUMN week_start_date SET NOT NULL;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE pay_periods
    ALTER COLUMN period_start_date SET NOT NULL,
    ALTER COLUMN period_end_date SET NOT NULL;
-- +goose StatementEnd
