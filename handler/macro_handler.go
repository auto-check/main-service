package handler

import (
	"context"
	"github.com/auto-check/main-service/usecase"
	mainpb "github.com/auto-check/protocol-buffer/golang/main"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MacroHandler struct {
	mainpb.MainServer
	su usecase.MacroUsecase
}

func NewMacroHandler(gserver *grpc.Server, us usecase.MacroUsecase) *MacroHandler {
	handler := &MacroHandler{
		su: us,
	}
	mainpb.RegisterMainServer(gserver, handler)
	return handler
}

func (sh *MacroHandler) CreateMacro(ctx context.Context, r *emptypb.Empty) (*emptypb.Empty, error) {
	id, _ := ctx.Value("student_id").(int64)

	err := sh.su.CreateMacroByStudentID(int64(id))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

