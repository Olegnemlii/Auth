package main

import (
	"context"
	"errors"
	"log"
	"net"

	desc "github.com/Olegnemlii/auth/pkg/auth_V1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedAuthV1Server
	users      map[int64]*desc.GetUserResponse
	lastUserID int64
}

func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	log.Printf("CreateUser Request received: %v", req)
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, errors.New("passwords do not match")
	}
	s.lastUserID++
	newUser := &desc.GetUserResponse{
		Id:        s.lastUserID,
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Role:      req.GetRole(),
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
	s.users[s.lastUserID] = newUser
	return &desc.CreateUserResponse{
		Id: s.lastUserID,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("GetUser Request received: %v", req)
	user, ok := s.users[req.GetId()]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *server) UpdatUser(ctx context.Context, req *desc.UpdateUserInfoRequest) (*desc.UpdateUserInfoResponse, error) {
	log.Printf("UpdateUser Request received: %v", req)
	user, ok := s.users[req.GetId()]
	if !ok {
		return nil, errors.New("user not found")
	}
	if req.Name != nil {
		user.Name = req.GetName().GetValue()
	}
	if req.Email != nil {
		user.Email = req.GetEmail().GetValue()
	}
	user.UpdatedAt = timestamppb.Now()
	return &desc.UpdateUserInfoResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*desc.DeleteUserResponse, error) {
	log.Printf("DeleteUser Request received: %v", req)
	_, ok := s.users[req.GetId()]
	if !ok {
		return nil, errors.New("user not found")
	}
	delete(s.users, req.GetId())
	return &desc.DeleteUserResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func main() {
	// 1. Настройка лисенера
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Создание gRPC сервера
	grpcServer := grpc.NewServer()
	server := &server{
		users:      make(map[int64]*desc.GetUserResponse),
		lastUserID: 0,
	}

	// 3. Регистрация сервиса
	desc.RegisterAuthV1Server(grpcServer, server)
	log.Println("Starting gRPC server on port 50052")

	// 4. Запуск сервера
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
