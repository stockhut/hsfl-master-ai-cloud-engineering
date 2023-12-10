package grpc_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "GRPC ", log.LstdFlags|log.Lmsgprefix)

type GrpcServer struct {
	auth_proto.UnimplementedAuthenticationServer
	repo repository.AccountRepository
}

func New(repo repository.AccountRepository) *GrpcServer {
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
	name := request.GetName()

	acc, err := s.repo.FindAccount(name)
	if err != nil {
		if errors.Is(err, repository.ErrAccountNotFound) {
			return nil, auth_proto.ErrAccountNotFound
		}
		return nil, err
	}

	response := auth_proto.AccountResponseFromModel(acc)
	return response, nil
}
