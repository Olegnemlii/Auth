package chat

import (
	"context"
	"log"

	"github.com/Olegnemlii/chat-server/internal/converter"
	desc "github.com/Olegnemlii/chat-server/pkg/chat_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	id, err := i.chatService.CreateChat(ctx, converter.ToChatInfoFromDesc(req.Get()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return &desc.CreateChatResponse{
		Id: id,
	}, nil
}
