package auth_proto

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC error definitions
// see https://jbrandhorst.com/post/grpc-errors/

var ErrAccountNotFound = status.Error(codes.NotFound, "no such account")
