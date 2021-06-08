package repository

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

type macroRepository struct {
	DB *sql.DB
}

func NewMacroRepository(db *sql.DB) MacroRepository {
	return &macroRepository{
		DB: db,
	}
}

func (mr *macroRepository) CreateMacroByStudentID(id int64) error{
	query := `INSERT INTO macro SET student_id=?, onoff=?`
	stmt, err := mr.DB.Prepare(query)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return err
	}

	_, err = stmt.Exec(id, false)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return err
	}

	return nil
}
