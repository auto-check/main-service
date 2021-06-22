package mocks

import (
	"context"
	gomock "github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockMainClient struct {
	ctrl *gomock.Controller
	recorder *mockMainClientRecorder
}

type mockMainClientRecorder struct {
	mock *MockMainClient
}

func NewMockMainClient(ctrl *gomock.Controller) *MockMainClient{
	mock := &MockMainClient{ctrl: ctrl}
	mock.recorder = &mockMainClientRecorder{mock}
	return mock
}

func (m *MockMainClient) EXPECT() *mockMainClientRecorder {
	return m.recorder
}

func (m *MockMainClient) CreateMacro(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	s := []interface{}{ctx, empty}
	ret := m.ctrl.Call(m, "CreateMacro", s...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder) CreateMacro(arg0, arg1 interface{}) *gomock.Call {
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateMacro", s...)
}