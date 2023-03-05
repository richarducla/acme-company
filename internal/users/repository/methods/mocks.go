package methods

import (
	"errors"
	"regexp"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

const (
	insertUserQuery = "INSERT INTO \"users\""
	selectUserQuery = "SELECT * FROM \"users\""
)

func FailInInsertUserMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(insertUserQuery)).WillReturnError(errors.New("fail in bd"))
	mock.ExpectRollback()
}

func GivenInserUserMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(insertUserQuery)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
}

func FailInGetUserMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(selectUserQuery)).
		WillReturnRows(sqlmock.NewRows(nil))
	mock.ExpectRollback()
}

func GivenGetUserMock(mock sqlmock.Sqlmock) {
	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "email", "status", "created_at", "uptated_at"}).
		AddRow(1, "richard", "$ji4123%42123", "r@test.com", true, time.Now(), time.Now())
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(selectUserQuery)).
		WillReturnRows(rows)
	mock.ExpectCommit()
}
