package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Olegnemlii/chat-server/pkg/chat_v1"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := desc.NewChatV1Client(conn)

	createChat(client)

	deleteChat(client)

	sendMessage(client)

}

func createChat(client desc.ChatV1Client) {
	log.Println("CreateChat")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &desc.CreateChatRequest{
		Usernames: "user1,user2",
	}
	resp, err := client.CreateChat(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call CreateChat: %v", err)
	}
	log.Printf("CreateChat Response: ID=%d\n", resp.GetId())
}

func deleteChat(client desc.ChatV1Client) {
	log.Println("DeleteChat")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &desc.DeleteChatRequest{
		Id: 1,
	}
	resp, err := client.DeleteChat(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call DeleteChat: %v", err)
	}
	log.Printf("DeleteChat Response: %v\n", resp)
}

func sendMessage(client desc.ChatV1Client) {
	log.Println("SendMessage")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &desc.SendMessageRequest{
		From:      "user1",
		Text:      "Hello, how are you?",
		Timestamp: timestamppb.Now(),
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call SendMessage: %v", err)
	}

	log.Printf("SendMessage Response: %v\n", resp)
}
