package main

import (
	"database/sql"
	"fmt"
	"github.com/auto-check/common-module/middleware"
	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"main/handler"
	"main/repository"
	"main/usecase"
	"net"
	"net/url"
	"os"
)

func init() {
	err := 	godotenv.Load()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	connection = fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, connection)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	sr := repository.NewMacroRepository(dbConn)
	su := usecase.NewStudentUsecase(sr)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.Authenticator))))

	handler.NewMacroHandler(server, su)

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("start gRPC server on %s port\n", os.Getenv("PORT"))
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}