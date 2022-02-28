package main

import (
	"github.com/gin-gonic/gin"
	handler2 "github.com/shinshin8/golang-grpc-client/handler"
	service2 "github.com/shinshin8/golang-grpc-client/service"
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("GRPC_PORT", "8081")
	os.Setenv("PORT", "8082")

	conn, err := grpc.Dial(os.Getenv("HOST")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcService := pb.NewGrpcServiceClient(conn)

	service := service2.NewService(grpcService)

	handler := handler2.NewHandler(service)

	gin.SetMode(gin.DebugMode)
	engin := gin.Default()
	engin = (handler2.SetRoute(handler))(engin)

	if err := engin.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
