package grpc_server

import (
	"context"
	mock_accounts "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/proto"
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
			testName string
			accName  string
			acc      model.Account
			err      error
		}

		testCases := []testCase{
			{
				testName: "account found",
				acc:      bobAccount,
				accName:  bobAccount.Name,
				err:      nil,
			},
			{
				testName: "account not found",
				acc:      model.Account{},
				accName:  "no-account-for-this-name",
				err:      repository.ErrAccountNotFound,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				mockCtrl := gomock.NewController(t)
				mockRepo := mock_accounts.NewMockAccountRepository(mockCtrl)

				mockRepo.EXPECT().FindAccount(tc.accName).Return(&tc.acc, tc.err).Times(1)

				grpcServer := New(mockRepo)

				acc, err := grpcServer.GetAccount(context.Background(), &proto.GetAccountRequest{Name: tc.accName})

				assert.Equal(t, tc.err, err)

				if tc.err == nil {
					assert.Equal(t, tc.acc.Name, acc.Name)
					assert.Equal(t, tc.acc.Email, acc.Email)
				}
			})
		}

	})

}
