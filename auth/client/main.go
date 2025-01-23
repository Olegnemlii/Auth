package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/Olegnemlii/auth/pkg/auth_V1"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := desc.NewAuthV1Client(conn)

	createUser(client)

	getUser(client)

	updateUser(client)

	deleteUser(client)
}

func createUser(client desc.AuthV1Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateUser(ctx, &desc.CreateUserRequest{
		Name:            "Test User",
		Email:           "test@example.com",
		Password:        "password123",
		PasswordConfirm: "password123",
		Role:            desc.UserRole_USER,
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("User created with ID: %d\n", resp.GetId())
}

func getUser(client desc.AuthV1Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &desc.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	fmt.Printf("User: %+v\n", resp)
}

func updateUser(client desc.AuthV1Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.UpdatUser(ctx, &desc.UpdateUserInfoRequest{
		Id:    1,
		Name:  wrapperspb.String("Updated Name"),
		Email: wrapperspb.String("updated@example.com"),
	})
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	fmt.Println("User updated successfully")
}

func deleteUser(client desc.AuthV1Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.DeleteUser(ctx, &desc.DeleteUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("User deleted successfully")
}
