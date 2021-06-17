package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MacroRepository struct {
	mock.Mock
}

func (_m *MacroRepository) Store(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(id int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//func (_m *StudentRepository) GetByGcn(gcn string) (*model.Student, error){
//	ret := _m.Called(gcn)
//
//	var r0 *model.Student
//	if rf, ok := ret.Get(0).(func(string) *model.Student); ok {
//		r0 = rf(gcn)
//	} else {
//		r0 = ret.Get(0).(*model.Student)
//	}
//
//	var r1 error
//	if rf, ok := ret.Get(1).(func(string) error); ok {
//		r1 = rf(gcn)
//	} else {
//		r1 = ret.Error(1)
//	}
//
//	return r0, r1
//}
//
//
//
//


