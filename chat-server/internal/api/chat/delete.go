package chat

import (
	"context"
	"log"

	"github.com/Olegnemlii/chat-server/internal/converter"
	desc "github.com/Olegnemlii/chat-server/pkg/chat_v1"
)

func (i *Implementation) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*desc.DeleteChatResponse, error) {
	chatObj, err := i.chatService.DeleteChat(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", chatObj.ID, chatObj.Info.Title, chatObj.Info.Content, chatObj.CreatedAt, chatObj.UpdatedAt)

	return &desc.DeleteChatResponse{
		Chat: converter.ToChatFromService(chatObj),
	}, nil
}
