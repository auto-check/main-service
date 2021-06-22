package client

import (
	mainpb "github.com/auto-check/protocol-buffer/golang/main"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"sync"
)

func init() {
	err := 	godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	once sync.Once
	cli mainpb.MainClient
)

func GetMainClient() mainpb.MainClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			"127.0.0.1:"+os.Getenv("PORT"),
			grpc.WithInsecure(),
			grpc.WithBlock())

		cli = mainpb.NewMainClient(conn)
	})

	return cli
}

