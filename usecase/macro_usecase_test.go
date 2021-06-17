package usecase

import (
	"github.com/stretchr/testify/assert"
	"main/repository/mocks"
	"testing"
)

func TestCreateMacroByStudentID(t *testing.T) {
	mockMacroRepo := new(mocks.MacroRepository)

	t.Run("success", func(t *testing.T){
		mockMacroRepo.On("Store", int64(1)).
			Return(nil).Once()

		mu := NewMacroUsecase(mockMacroRepo)
		err := mu.CreateMacroByStudentID(1)

		assert.NoError(t, err)
	})

}
