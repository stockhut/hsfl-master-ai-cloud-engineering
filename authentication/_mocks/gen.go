package mocks

//go:generate mockgen -source ../accounts/repository.go -destination=./repository_mocks/repository_mocks.go
//go:generate mockgen -source ../pwhash/pwhasher.go -destination=./pwhash_mocks/pwhash_mocks.go
//go:generate mockgen -source ../auth-proto/authentication_grpc.pb.go -destination=./mock-auth-proto/auth-proto_mocks.go
