package book

import "context"
проверка изменения ветки main
type Book struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
	Available bool   `json:"available"`
}

type Service interface {
	Create(ctx context.Context, b Book) (string, error)
	GetByID(ctx context.Context, id string) (Book, error)
	UpdateAuthor(ctx context.Context, id string, author string) error
}

type Repository interface {
	CreateBook(ctx context.Context, b Book) error
	GetBookByID(ctx context.Context, id string) (Book, error)
	UpdateBookAuthor(ctx context.Context, id string, author string) error
}
