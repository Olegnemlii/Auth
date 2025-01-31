package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Olegnemlii/chat-server/internal/client/db"
	"github.com/Olegnemlii/chat-server/internal/model"
	"github.com/Olegnemlii/chat-server/internal/repository"
	"github.com/Olegnemlii/chat-server/internal/repository/chat/converter"
	modelRepo "github.com/Olegnemlii/chat-server/internal/repository/chat/model"
)

const (
	tableName = "chat"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) CreateChat(ctx context.Context, info *model.ChatInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		Values(info.Title, info.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) DeleteChat(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColumn, titleColumn, contentColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chat.ID, &chat.Info.Title, &chat.Info.Content, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}
