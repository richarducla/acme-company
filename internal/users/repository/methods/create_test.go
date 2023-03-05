package methods

import (
	"acme/internal/models"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_Create_WhenTheRecordIsSaved_Success(t *testing.T) {
	testCases := []struct {
		log      string
		input    models.User
		expected models.User
		failed   bool
		fnCalled []func(sqlmock.Sqlmock)
	}{
		{
			log: "Fail in create user",
			input: models.User{
				Username: "richard",
				Email:    "r@test.com",
				Password: "1234",
			},
			failed: true,
			fnCalled: []func(sqlmock.Sqlmock){
				FailInInsertUserMock,
			},
		},
		{
			log: "Fail in get user",
			input: models.User{
				Username: "richard",
				Email:    "r@test.com",
				Password: "1234",
			},
			failed: true,
			fnCalled: []func(sqlmock.Sqlmock){
				GivenInserUserMock,
				FailInGetUserMock,
			},
		},
		{
			log: "Created user sucessfully",
			input: models.User{
				Username: "richard",
				Email:    "r@test.com",
				Password: "1234",
			},
			failed: false,
			fnCalled: []func(sqlmock.Sqlmock){
				GivenInserUserMock,
				GivenGetUserMock,
			},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.log)

		w, err := newTestWrapper()
		if err != nil {
			t.Error(err)
		}
		w.mock.MatchExpectationsInOrder(false)

		for _, fn := range tc.fnCalled {
			fn(w.mock)
		}

		result, err := w.repository.Create(tc.input)
		log.Println(err)

		w.cleanTestWrapper()

		if tc.failed {
			assert.Error(t, err, "fail in bd")
		} else {
			assert.NotEmpty(t, result)
			assert.IsType(t, models.User{}, result)
			assert.Nil(t, err)
		}
	}
}
