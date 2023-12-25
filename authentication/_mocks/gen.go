package _mocks

//go:generate mockgen -source ../accounts/repository/repository.go -destination=./repository_mocks/repository_mocks.go
//go:generate mockgen -source ../pwhash/pwhasher.go -destination=./pwhash_mocks/pwhash_mocks.go
