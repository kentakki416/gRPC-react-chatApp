package main

import (
	"context"
	"log"

	"github.com/kentakki416/gRPC-react-chatApp/domain"
	pb "github.com/kentakki416/gRPC-react-chatApp/pb"
	"github.com/kentakki416/gRPC-react-chatApp/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	usecase.CreateMessageService
	usecase.GetMessageStreamService
}

func NewServer(createMessagesService usecase.CreateMessageService, getMessageStreamService usecase.GetMessageStreamService) *Server {
	return &Server{
		CreateMessageService:    createMessagesService,
		GetMessageStreamService: getMessageStreamService,
	}
}

func (s *Server) GetMessageStream(req *emptypb.Empty, server pb.ChatService_GetMessageStreamServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan domain.Message)

	go func() {
		if err := s.GetMessageStreamService.Handle(ctx, stream); err != nil {
			log.Println(err)
		}
	}()

	for {
		v := <-stream
		createdAd := timestamppb.New(v.CreatedAt)
		if err := server.Send(&pb.GetMessageStreamResponse{
			Message: &pb.Message{
				From:           v.From,
				MessageContent: v.MessageContent,
				CreatedAt:      createdAd,
			},
		}); err != nil {
			return err
		}
	}
}

func (s *Server) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	input := usecase.NewCreateMessageServiceInput(req.Message.From, req.Message.MessageContent, req.Message.CreatedAt.AsTime())
	if err := s.CreateMessageService.Handle(ctx, input); err != nil {
		return &pb.CreateMessageResponse{
			Result: err.Error(),
		}, err
	}
	return &pb.CreateMessageResponse{
		Result: "ok",
	}, nil
}
