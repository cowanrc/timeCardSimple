-- +goose Up
-- +goose StatementBegin
ALTER TABLE employees
    DROP COLUMN password;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------

-- +goose Down
-- +goose StatementBegin
ALTER TABLE employees
    ADD COLUMN password;
-- +goose StatementEnd

