package grpc_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "GRPC ", log.LstdFlags|log.Lmsgprefix)

type GrpcServer struct {
	auth_proto.UnimplementedAuthenticationServer
	repo accounts.Repository
}

func New(repo accounts.Repository) *GrpcServer {
	return &GrpcServer{
		repo: repo,
	}
}

func (s *GrpcServer) Serve(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Printf("Failed to listen on port %d: %s\n", port, err)
		return err
	}

	srv := grpc.NewServer()
	auth_proto.RegisterAuthenticationServer(srv, s)

	logger.Printf("listening on port %d\n", port)
	err = srv.Serve(listener)
	if err != nil {
		logger.Printf("Error: %s\n", err)
	}
	return err
}

func (s *GrpcServer) GetAccount(ctx context.Context, request *auth_proto.GetAccountRequest) (*auth_proto.GetAccountResponse, error) {
	logger.Printf("GetAccount: %s\n", request.Name)
	name := request.GetName()

	acc, err := s.repo.FindAccount(ctx, name)
	if err != nil {
		if errors.Is(err, accounts.ErrAccountNotFound) {
			return nil, auth_proto.ErrAccountNotFound
		}
		return nil, auth_proto.ErrInternal
	}

	response := auth_proto.AccountResponseFromModel(acc)
	return response, nil
}
