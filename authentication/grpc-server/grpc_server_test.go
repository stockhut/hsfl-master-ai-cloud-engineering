package grpc_server

import (
	"context"
	mock_repository "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/mock-repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGrpcServer_GetAccount(t *testing.T) {

	t.Run("returns repository response", func(t *testing.T) {

		bobAccount := model.Account{
			Name:     "bob",
			Email:    "bob@example.org",
			Password: "",
		}

		type testCase struct {
			testName    string
			accName     string
			acc         model.Account
			expectedErr error
			repoErr     error
		}

		testCases := []testCase{
			{
				testName:    "account found",
				acc:         bobAccount,
				accName:     bobAccount.Name,
				expectedErr: nil,
				repoErr:     nil,
			},
			{
				testName:    "account not found",
				acc:         model.Account{},
				accName:     "no-account-for-this-name",
				expectedErr: auth_proto.ErrAccountNotFound,
				repoErr:     repository.ErrAccountNotFound,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				mockCtrl := gomock.NewController(t)
				mockRepo := mock_repository.NewMockAccountRepository(mockCtrl)

				mockRepo.EXPECT().FindAccount(tc.accName).Return(&tc.acc, tc.repoErr).Times(1)

				grpcServer := New(mockRepo)

				acc, err := grpcServer.GetAccount(context.Background(), &auth_proto.GetAccountRequest{Name: tc.accName})

				assert.Equal(t, tc.expectedErr, err)

				if tc.expectedErr == nil {
					assert.Equal(t, tc.acc.Name, acc.Name)
					assert.Equal(t, tc.acc.Email, acc.Email)
				}
			})
		}

	})

}
