package service

import (
	"context"

	"github.com/Olegnemlii/chat-server/internal/model"
)

type ChatService interface {
	CreateChat(ctx context.Context, info *model.ChatInfo) (int64, error)
	DeleteChat(ctx context.Context, id int64) (*model.Chat, error)
}
