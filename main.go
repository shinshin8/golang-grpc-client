package main

import (
	"github.com/gin-gonic/gin"
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial(os.Getenv("HOST")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcService := pb.NewGrpcServiceClient(conn)

	service := NewService(grpcService)

	handler := NewHandler(service)

	gin.SetMode(gin.DebugMode)
	engin := gin.Default()
	engin = (SetRoute(handler))(engin)

	if err := engin.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
