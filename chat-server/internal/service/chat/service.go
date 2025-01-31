package chat

import (
	"github.com/Olegnemlii/chat-server/internal/client/db"
	"github.com/Olegnemlii/chat-server/internal/repository"
	"github.com/Olegnemlii/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
