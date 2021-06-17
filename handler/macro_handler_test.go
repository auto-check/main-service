package handler

import (
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/usecase/mocks"
	"testing"
)

func init() {
	err := 	godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateMacro(t *testing.T) {
	mockStudent := new(mocks.MacroUsecase)
	mockStudent.On("CreateMacroByStudentID", int64(1)).Return(nil)

	mh := NewMacroHandler(grpc.NewServer(), mockStudent)
	ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"student_id": []string{"1"}})
	log.Println(ctx.Value("student_id"))
	_, err := mh.CreateMacro(ctx, &emptypb.Empty{})

	assert.NoError(t, err)
}
