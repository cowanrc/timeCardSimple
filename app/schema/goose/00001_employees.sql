-- +goose Up
-- +goose StatementBegin

create table employees (
	id uuid not null,
	first_name character varying(256) default ''::character varying,
    last_name character varying(256) default ''::character varying,
	email citext not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null,
    password character varying(256),

    constraint "employees_pkey" primary key(id),
    constraint "employees_email_key" unique(email)
)
;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
drop table employees
;
-- +goose StatementEnd
