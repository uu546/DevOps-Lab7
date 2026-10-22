package repository

import (
	"context"

	"example.com/book-service/book"
)

type memoryRepo struct {
	books map[string]book.Book
}

// NewMemoryRepository создаёт in-memory хранилище книг.
func NewMemoryRepository() book.Repository {
	return &memoryRepo{
		books: make(map[string]book.Book),
	}
}

func (r *memoryRepo) CreateBook(ctx context.Context, b book.Book) error {
	r.books[b.ID] = b
	return nil
}

func (r *memoryRepo) GetBookByID(ctx context.Context, id string) (book.Book, error) {
	b, ok := r.books[id]
	if !ok {
		return book.Book{}, book.ErrBookNotFound
	}
	return b, nil
}

func (r *memoryRepo) UpdateBookAuthor(ctx context.Context, id string, author string) error {
	b, ok := r.books[id]
	if !ok {
		return book.ErrBookNotFound
	}
	b.Author = author
	r.books[id] = b
	return nil
}