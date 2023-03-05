package methods

import (
	"acme/internal/users/domain"
	"database/sql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testWrapper struct {
	db         *sql.DB
	mock       sqlmock.Sqlmock
	repository domain.Repository
}

func newTestWrapper() (*testWrapper, error) {
	var db *sql.DB
	var err error

	// db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // use equal matcher
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &testWrapper{
		db:         db,
		mock:       mock,
		repository: NewRepository(gdb),
	}, nil
}

func (t *testWrapper) cleanTestWrapper() {
	t.db.Close()
}
