package repository

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := regexp.QuoteMeta(`INSERT INTO macro SET student_id=?, onoff=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1, false).WillReturnResult(sqlmock.NewResult(1, 1))

	sr := NewMacroRepository(db)

	err = sr.Store(1)
	assert.NoError(t, err)
}
