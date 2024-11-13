package sqlrepo

import (
	"context"
	"database/sql"
	"fmt"
	"text/template"
	"time"

	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/infra/trace"
)

type QueryExecerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row

	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type QueryExecerContextUsingTemplate interface {
	QueryExecerContext

	QueryRowContextUsingTemplate(ctx context.Context, t *template.Template, data any) (*sql.Row, error)

	QueryContextUsingTemplate(ctx context.Context, t *template.Template, data any) (*sql.Rows, error)

	ExecContextUsingTemplate(ctx context.Context, t *template.Template, data any) (sql.Result, error)

	TemplateQueryer
}

func queryRowContextUsingTemplate(
	ctx context.Context,
	qecut QueryExecerContextUsingTemplate,
	t *template.Template,
	data any,
) (*sql.Row, error) {
	ctx, query, args, err := ExecuteTemplateForQueryArgs(ctx, qecut, t, data)
	if err != nil {
		return nil, err
	}

	return qecut.QueryRowContext(ctx, query, args...), nil
}

func queryContextUsingTemplate(
	ctx context.Context,
	qecut QueryExecerContextUsingTemplate,
	t *template.Template,
	data any,
) (*sql.Rows, error) {
	ctx, query, args, err := ExecuteTemplateForQueryArgs(ctx, qecut, t, data)
	if err != nil {
		return nil, err
	}

	return qecut.QueryContext(ctx, query, args...)
}

func execContextUsingTemplate(
	ctx context.Context,
	qecut QueryExecerContextUsingTemplate,
	t *template.Template,
	data any,
) (sql.Result, error) {
	ctx, query, args, err := ExecuteTemplateForQueryArgs(ctx, qecut, t, data)
	if err != nil {
		return nil, err
	}

	return qecut.ExecContext(ctx, query, args...)
}

func execContextWithDialectErrorHandling(
	ctx context.Context,
	qec QueryExecerContext,
	ed ErrorDialect,
	query string,
	args ...any,
) (sql.Result, error) {
	result, err := qec.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, ed.NormalizeSQLError(err)
	}
	return result, nil
}

func ScanQueryRowContextUsingTemplate[A any](
	ctx context.Context,
	qecut QueryExecerContextUsingTemplate,
	t *template.Template,
	data any,
	sf ScanFunc[A],
) (a A, err error) {
	ctx, span := trace.StartSpan(ctx, "sqlrepov2.ScanQueryRowContextUsingTemplate")
	defer span.Finish()

	var row *sql.Row
	row, err = queryRowContextUsingTemplate(ctx, qecut, t, data)
	if err != nil {
		return
	}
	return ScanRow(row, sf)
}

func ScanQueryRowContext[A any](
	ctx context.Context,
	qec QueryExecerContext,
	query string,
	args []any,
	sf ScanFunc[A],
) (A, error) {
	ctx, span := trace.StartSpan(ctx, "sqlrepov2.ScanQueryRowContext")
	defer span.Finish()

	return ScanRow(qec.QueryRowContext(ctx, query, args...), sf)
}

func ScanQueryContextUsingTemplate[A any](
	ctx context.Context,
	qecut QueryExecerContextUsingTemplate,
	t *template.Template,
	data any,
	sf ScanFunc[A],
) ([]A, error) {
	ctx, span := trace.StartSpan(ctx, "sqlrepov2.ScanQueryContextUsingTemplate")
	defer span.Finish()

	rows, err := queryContextUsingTemplate(ctx, qecut, t, data)
	if err != nil {
		return nil, err
	}
	return ScanRows(rows, sf)
}

func ScanQueryContext[A any](
	ctx context.Context,
	qec QueryExecerContext,
	query string,
	args []any,
	sf ScanFunc[A],
) ([]A, error) {
	ctx, span := trace.StartSpan(ctx, "sqlrepov2.ScanQueryContext")
	defer span.Finish()

	rows, err := qec.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return ScanRows(rows, sf)
}

func ScanRows[A any](rows *sql.Rows, sf ScanFunc[A]) ([]A, error) {
	var result []A
	var err error
	for rows.Next() {
		var a A
		a, err = ScanRow(rows, sf)
		if err != nil {
			break
		}
		result = append(result, a)
	}
	errClose := rows.Close()
	if err != nil {
		return nil, err
	}
	if errClose != nil {
		return nil, errClose
	}
	return result, nil
}

func ScanRow[A any](rs RowScanner, sf ScanFunc[A]) (a A, err error) {
	return sf(rs)
}

type RowScanner interface {
	Scan(dst ...any) error
}

type ScanFunc[A any] func(RowScanner) (A, error)

type scanID struct {
	result *id.ID
}

func ScanIntoID(id *id.ID) *scanID {
	return &scanID{result: id}
}

func ScanIntoTime(t **time.Time) ScanFunc[*time.Time] {
	return func(rs RowScanner) (*time.Time, error) {
		var scannedTime sql.NullTime
		if err := rs.Scan(&scannedTime); err != nil {
			return nil, err
		}
		if !scannedTime.Valid {
			return nil, nil
		}
		*t = &scannedTime.Time
		return *t, nil
	}
}

func ScanIntoFloat64(f **float64) ScanFunc[*float64] {
	return func(rs RowScanner) (*float64, error) {
		var scannedFloat sql.NullFloat64
		if err := rs.Scan(&scannedFloat); err != nil {
			return nil, err
		}
		if !scannedFloat.Valid {
			return nil, nil
		}
		*f = &scannedFloat.Float64
		return *f, nil
	}
}

func (sid *scanID) Scan(src any) error {
	var err error

	switch srcType := src.(type) {
	case string:
		domainID, err := id.ParseString(srcType)
		if err == nil {
			*sid.result = domainID
			return nil
		}

	case []byte:
		domainID, err := id.ParseBytes(srcType)
		if err == nil {
			*sid.result = domainID
			return nil
		}

	default:
		err = fmt.Errorf("sqlrepo: cannot scan type %T in id.ID", src)
	}

	return err
}
