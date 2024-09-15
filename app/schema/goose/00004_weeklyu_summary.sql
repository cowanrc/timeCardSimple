-- +goose Up
-- +goose StatementBegin
create table weekly_summary (
    id uuid not null,
    employee_id uuid not null,
    week_start_date date not null,
    days_worked integer not null default 0,
	
    constraint "weekly_summary_pkey" primary key(id),
    constraint "weekly_summary_employee_id_fkey" foreign key(employee_id) references employees(id) on delete cascade
);
-- +goose StatementEnd

-- ------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
drop table weekly_summary
;
-- +goose StatementEnd
