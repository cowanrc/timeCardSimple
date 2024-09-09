-- +goose Up
-- +goose StatementBegin
create extension if
not
exists "citext"
;
-- +goose StatementEnd
-- ------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd

