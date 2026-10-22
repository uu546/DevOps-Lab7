package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"example.com/book-service/book"
)

// ----- DTO -----

type CreateRequest struct {
	Book book.Book
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

type GetByIDRequest struct {
	ID string `json:"id"`
}

type GetByIDResponse struct {
	Book book.Book `json:"book"`
	Err  error     `json:"error,omitempty"`
}

type UpdateAuthorRequest struct {
	ID     string `json:"id"`
	Author string `json:"author"`
}

type UpdateAuthorResponse struct {
	Err error `json:"error,omitempty"`
}

// ----- Endpoints -----

type Endpoints struct {
	Create       endpoint.Endpoint
	GetByID      endpoint.Endpoint
	UpdateAuthor endpoint.Endpoint
}

func MakeEndpoints(s book.Service) Endpoints {
	return Endpoints{
		Create:       makeCreateEndpoint(s),
		GetByID:      makeGetByIDEndpoint(s),
		UpdateAuthor: makeUpdateAuthorEndpoint(s),
	}
}

func makeCreateEndpoint(s book.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Book)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetByIDEndpoint(s book.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		b, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{Book: b, Err: err}, nil
	}
}

func makeUpdateAuthorEndpoint(s book.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAuthorRequest)
		err := s.UpdateAuthor(ctx, req.ID, req.Author)
		return UpdateAuthorResponse{Err: err}, nil
	}
}