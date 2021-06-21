package usecase

import (
	"github.com/auto-check/main-service/repository"
)

type macroUsecase struct {
	mr repository.MacroRepository
}

func NewMacroUsecase(mr repository.MacroRepository) MacroUsecase{
	return &macroUsecase{
		mr: mr,
	}
}

func (mu *macroUsecase) CreateMacroByStudentID(id int64) error {
	err := mu.mr.Store(id)
	if err != nil {
		return err
	}

	return nil
}