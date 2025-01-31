package converter

import (
	"github.com/Olegnemlii/chat-server/internal/model"
	modelRepo "github.com/Olegnemlii/chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        chat.ID,
		Info:      ToChatInfoFromRepo(chat.Info),
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}

func ToChatInfoFromRepo(info modelRepo.ChatInfo) model.ChatInfo {
	return model.ChatInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
