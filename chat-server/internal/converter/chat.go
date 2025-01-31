package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Olegnemlii/chat-server/internal/model"
	desc "github.com/Olegnemlii/chat-server/pkg/chat_v1"
)

func ToChatFromService(chat *model.Chat) *desc.Chat {
	var updatedAt *timestamppb.Timestamp
	if chat.UpdatedAt.Valid {
		updatedAt = timestamppb.New(chat.UpdatedAt.Time)
	}

	return &desc.Chat{
		Id:        chat.ID,
		Info:      ToChatInfoFromService(chat.Info),
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToChatInfoFromService(info model.ChatInfo) *desc.ChatInfo {
	return &desc.ChatInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToChatInfoFromDesc(info *desc.ChatInfo) *model.ChatInfo {
	return &model.ChatInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
