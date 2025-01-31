package chat

import (
	"context"

	"github.com/Olegnemlii/chat-server/internal/model"
)

func (s *serv) DeleteChat(ctx context.Context, id int64) (*model.Chat, error) {
	note, err := s.chatRepository.DeleteChat(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
