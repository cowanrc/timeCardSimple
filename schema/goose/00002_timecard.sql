-- +goose Up
-- +goose StatementBegin
create table timecard (
	id uuid not null,
	employee_id uuid not null,
	start_time timestamp without time zone not null,
	end_time timestamp without time zone not null,
	created_at timestamp without time zone not null,
	updated_at timestamp without time zone not null,
	
	
	constraint "timecard_pkey" primary key(id),
    constraint "timecard_employee_id_fkey" foreign key(employee_id) references employees(id) on delete cascade
)
-- +goose StatementEnd

-- ------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
drop table timecard
;
-- +goose StatementEnd
