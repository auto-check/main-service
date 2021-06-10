package repository

type MacroRepository interface {
	Store(id int64) error
}