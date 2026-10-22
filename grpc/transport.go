package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"example.com/book-service/book"
	"example.com/book-service/pb"
	"example.com/book-service/transport"
)

type grpcServer struct {
	pb.UnimplementedBookServiceServer

	create       grpctransport.Handler
	getByID      grpctransport.Handler
	updateAuthor grpctransport.Handler
}

func NewGRPCServer(endpoints transport.Endpoints) pb.BookServiceServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.Create,
			decodeCreateRequest,
			encodeCreateResponse,
		),
		getByID: grpctransport.NewServer(
			endpoints.GetByID,
			decodeGetByIDRequest,
			encodeGetByIDResponse,
		),
		updateAuthor: grpctransport.NewServer(
			endpoints.UpdateAuthor,
			decodeUpdateAuthorRequest,
			encodeUpdateAuthorResponse,
		),
	}
}

func (s *grpcServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateResponse), nil
}

func (s *grpcServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	_, resp, err := s.getByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetByIDResponse), nil
}

func (s *grpcServer) UpdateAuthor(ctx context.Context, req *pb.UpdateAuthorRequest) (*pb.UpdateAuthorResponse, error) {
	_, resp, err := s.updateAuthor.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UpdateAuthorResponse), nil
}

func decodeCreateRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.CreateRequest)
	return transport.CreateRequest{
		Book: book.Book{
			Title:  r.Title,
			Author: r.Author,
			Year:   int(r.Year),
			ISBN:   r.Isbn,
		},
	}, nil
}

func encodeCreateResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(transport.CreateResponse)
	errStr := ""
	if r.Err != nil {
		errStr = r.Err.Error()
	}
	return &pb.CreateResponse{Id: r.ID, Error: errStr}, nil   // ← Id, не ID
}

func decodeGetByIDRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetByIDRequest)
	return transport.GetByIDRequest{ID: r.Id}, nil   // ← r.Id
}

func encodeGetByIDResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(transport.GetByIDResponse)
	errStr := ""
	if r.Err != nil {
		errStr = r.Err.Error()
	}
	return &pb.GetByIDResponse{
		Book: &pb.Book{
			Id:        r.Book.ID,
			Title:     r.Book.Title,
			Author:    r.Book.Author,
			Year:      int32(r.Book.Year),
			Isbn:      r.Book.ISBN,      
			Available: r.Book.Available,
		},
		Error: errStr,
	}, nil
}

func decodeUpdateAuthorRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.UpdateAuthorRequest)
	return transport.UpdateAuthorRequest{ID: r.Id, Author: r.Author}, nil   // ← r.Id
}

func encodeUpdateAuthorResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(transport.UpdateAuthorResponse)
	errStr := ""
	if r.Err != nil {
		errStr = r.Err.Error()
	}
	return &pb.UpdateAuthorResponse{Error: errStr}, nil
}