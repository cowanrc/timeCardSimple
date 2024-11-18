package employeesql

import (
	"testing"
	"timeCardSimple/app/infra/storage/domainsql/employeesql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setupMockDB(t *testing.T) (*Repo, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := employeesql.New(db)
	return repo, mock
}
