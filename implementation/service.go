package implementation

import (
	"context"

	"github.com/google/uuid"

	"example.com/book-service/book"
)

type service struct {
	repository book.Repository
}

// NewService создаёт сервис библиотеки.
func NewService(repo book.Repository) book.Service {
	return &service{repository: repo}
}

func (s *service) Create(ctx context.Context, b book.Book) (string, error) {
	id := uuid.New().String()
	b.ID = id
	b.Available = true

	if err := s.repository.CreateBook(ctx, b); err != nil {
		return "", book.ErrCmdRepository
	}
	return id, nil
}

func (s *service) GetByID(ctx context.Context, id string) (book.Book, error) {
	b, err := s.repository.GetBookByID(ctx, id)
	if err != nil {
		return book.Book{}, book.ErrBookNotFound
	}
	return b, nil
}

func (s *service) UpdateAuthor(ctx context.Context, id string, author string) error {
	if err := s.repository.UpdateBookAuthor(ctx, id, author); err != nil {
		return book.ErrCmdRepository
	}
	return nil
}