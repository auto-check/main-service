package handler

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	mainpb "main/protocol-buffer/golang/main"
	"main/usecase"
	"runtime/debug"
	"strconv"
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
	md, _ := metadata.FromIncomingContext(ctx)
	id, err := strconv.Atoi(md.Get("student_id")[0])
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return &emptypb.Empty{}, err
	}

	err = sh.su.CreateMacroByStudentID(int64(id))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
