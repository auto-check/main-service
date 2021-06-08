package repository

type MacroRepository interface {
	CreateMacroByStudentID(id int64) error
}