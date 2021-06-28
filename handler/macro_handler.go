package handler

import (
	"context"
	"github.com/auto-check/main-service/usecase"
	authpb "github.com/auto-check/protocol-buffer/golang/auth"
	mainpb "github.com/auto-check/protocol-buffer/golang/main"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MacroHandler struct {
	mainpb.MainServer
	su usecase.MacroUsecase
	ac authpb.AuthClient
}

func NewMacroHandler(gserver *grpc.Server, us usecase.MacroUsecase, ac authpb.AuthClient) *MacroHandler {
	handler := &MacroHandler{
		su: us,
		ac: ac,
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

func (sh *MacroHandler) GetMacroStatus(ctx context.Context, r *emptypb.Empty) (*mainpb.GetMacroStatusResponse, error) {
	id, _ := ctx.Value("student_id").(int64)
	log.Println(id)
}

