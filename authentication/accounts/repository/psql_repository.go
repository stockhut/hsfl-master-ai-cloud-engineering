package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository/_sqlc"
)

type PsqlRepository struct {
	queries *sqlc.Queries
}

func NewPsqlRepository(connectionString string) (*PsqlRepository, error) {
	ctx := context.Background()
	// TODO: sslmode=disable should not be default, but it's too convenient atm
	db, err := sql.Open("pgx", connectionString+"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, Ddl); err != nil {
		fmt.Printf("Failed to create tables: %s\n", err)
		return nil, err
	}

	queries := sqlc.New(db)

	return &PsqlRepository{
		queries: queries,
	}, nil
}

func (repo *PsqlRepository) CreateAccount(ctx context.Context, acc model.Account) error {
	_, err := repo.queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Name:         acc.Name,
		Email:        acc.Email,
		Passwordhash: acc.PasswordHash,
	})

	return err
}

func (repo *PsqlRepository) CheckDuplicate(ctx context.Context, acc model.Account) (AccountInfoDuplicate, error) {

	_, err := repo.queries.GetAccount(ctx, acc.Name)
	if err == nil {
		return DUPLICATE_NAME, nil
	} else if errors.Is(err, sql.ErrNoRows) == false {
		return UNDEFINED, err

	}

	_, err = repo.queries.GetAccountByEmail(ctx, acc.Email)
	if err == nil {
		return DUPLICATE_EMAIL, nil
	} else if errors.Is(err, sql.ErrNoRows) == false {
		return UNDEFINED, err

	}

	return NO_DUPLICATES, nil
}

func (repo *PsqlRepository) FindAccount(ctx context.Context, name string) (*model.Account, error) {

	acc, err := repo.queries.GetAccount(ctx, name)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Account{
		Name:         acc.Name,
		Email:        acc.Email,
		PasswordHash: acc.Passwordhash,
	}, nil
}
