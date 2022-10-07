package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/kentakki416/gRPC-react-chatApp/pb"

	"github.com/kentakki416/gRPC-react-chatApp/infra"
	"github.com/kentakki416/gRPC-react-chatApp/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = 8080
)

func main() {
	ctx := context.Background()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	register(ctx, s)

	// サーバーリフレクション
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server, port: %d", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

func register(ctx context.Context, s *grpc.Server) {
	repo := infra.NewLocalMessageRepositoryImpl()
	createMessageService := usecase.NewCreateMessageService(repo)
	getMessageService := usecase.NewGetMessageStreamService(repo)
	pb.RegisterChatServiceServer(s, NewServer(*createMessageService, *getMessageService))
}
