-- +goose Up
-- +goose StatementBegin

alter table weekly_summary
    add column total_hours numeric not null default 0;

create table pay_periods (
    id uuid not null,
    employee_id uuid not null,
    period_start_date date not null,
    period_end_date date not null,
    total_days_worked integer not null default 0,
    total_hours numeric not null default 0,

    constraint "pay_periods_pkey" primary key(id),
    constraint "pay_periods_employee_id_fkey" foreign key(employee_id) references employees(id) on delete cascade,
    unique (employee_id, period_start_date, period_end_date)
);
-- +goose StatementEnd

-- ------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
alter table weekly_summary drop column total_hours;
drop table pay_periods;
-- +goose StatementEnd
