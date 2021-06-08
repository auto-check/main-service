package usecase

import (
	"main/repository"
)

type macroUsecase struct {
	mr repository.MacroRepository
}

func NewStudentUsecase(mr repository.MacroRepository) MacroUsecase{
	return &macroUsecase{
		mr: mr,
	}
}

func (mu *macroUsecase) CreateMacroByStudentID(id int64) error {
	err := mu.mr.CreateMacroByStudentID(id)
	if err != nil {
		return err
	}

	return nil
}