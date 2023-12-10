package _mocks

//go:generate mockgen -source ../accounts/repository/repository.go -destination=./mock-repository/repository_mocks.go
//go:generate mockgen -source ../auth-proto/authentication_grpc.pb.go -destination=./mock-auth-proto/auth-proto_mocks.go
