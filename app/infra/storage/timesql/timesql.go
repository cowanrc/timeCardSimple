package timesql

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ScanSQLTimeError struct {
	src interface{}
}

func (e *ScanSQLTimeError) Error() string {
	return fmt.Sprintf("sqlrepo: cannot scan type %T into SQLTime", e.src)
}

type SQLTime struct {
	//time is the underlying Time.
	//
	//It must be rounded and in UTC.
	time time.Time
}

func Now() SQLTime {
	return NewSQLTime(time.Now())
}

func NewSQLTime(t time.Time) SQLTime {
	return SQLTime{
		time: t.Round(time.Microsecond).In(time.UTC),
	}
}

func (t SQLTime) Value() (driver.Value, error) {
	return t.time, nil
}

func (t *SQLTime) Scan(src interface{}) error {
	if srcTime, ok := src.(time.Time); ok {
		t.time = srcTime.Round(time.Microsecond).UTC()
		return nil
	}

	return &ScanSQLTimeError{src: src}
}

func (t SQLTime) Equal(other SQLTime) bool {
	return t.time.Equal(other.time)
}

func (t SQLTime) Domain() time.Time {
	return t.time
}

type NullSQLTime struct {
	SQLTime SQLTime

	Valid bool
}

func NewNullSQLTime(t *time.Time) NullSQLTime {
	result := NullSQLTime{
		Valid: t != nil,
	}
	if result.Valid {
		result.SQLTime = NewSQLTime(*t)
	}
	return result
}

func (t NullSQLTime) Value() (driver.Value, error) {
	if t.Valid {
		return t.SQLTime.Value()
	}
	return nil, nil
}

func (t *NullSQLTime) Scan(src interface{}) error {
	if src == nil {
		t.Valid = false
		return nil
	}

	err := t.SQLTime.Scan(src)
	if err != nil {
		t.Valid = false
		return err
	}

	t.Valid = true
	return nil
}

func (t NullSQLTime) Domain() *time.Time {
	if t.Valid {
		domain := t.SQLTime.Domain()
		return &domain
	}
	return nil
}
