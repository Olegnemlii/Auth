package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/Olegnemlii/chat-server/pkg/chat_v1"

	"github.com/brianvoe/gofakeit"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

// Get ...
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Chat usernames: %s", req.GetUsernames())

	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*desc.DeleteChatResponse, error) {

	log.Printf("DeleteChatRequest received: %v", req)

	if req.GetId() == 0 {
		return nil,
			fmt.Errorf("cannot delete chat with id = 0")
	}

	return &desc.DeleteChatResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	log.Printf("SendMessage Request received: %v", req)
	from := req.GetFrom()
	text := req.GetText()
	timestamp := req.GetTimestamp()

	log.Printf("Message: From=%s, Text=%s, Timestamp=%v\n", from, text, timestamp)

	return &desc.SendMessageResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
