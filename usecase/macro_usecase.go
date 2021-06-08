package usecase

type MacroUsecase interface {
	CreateMacroByStudentID(id int64) error
}