package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MacroUsecase struct {
	mock.Mock
}

func (_m *MacroUsecase) CreateMacroByStudentID(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(id int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
