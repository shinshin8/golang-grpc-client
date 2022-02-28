package service

import (
	"context"
	"fmt"
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"strconv"
)

type Service interface {
	FindEmployee(id string) (*pb.FindEmployeeResponse, error)
	ListEmployee() (*pb.ListEmployeeResponse, error)
}

type service struct {
	client pb.GrpcServiceClient
}

func NewService(client pb.GrpcServiceClient) Service {
	return &service{client: client}
}

func (s *service) FindEmployee(id string) (*pb.FindEmployeeResponse, error) {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id: %w", err)
	}

	res, err := s.client.FindEmployee(context.Background(), &pb.FindEmployeeRequest{ID: int32(convID)})
	if err != nil {
		return nil, fmt.Errorf("failed to call FindEmployee: %w", err)
	}
	return res, nil
}

func (s *service) ListEmployee() (*pb.ListEmployeeResponse, error) {
	res, err := s.client.ListEmployee(context.Background(), &pb.ListEmployeeRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to call ListEmployee: %w", err)
	}
	return res, nil
}
