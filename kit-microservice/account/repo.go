package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoError = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{db, log.With(logger, "repo", "sql")}
}

func (r repo) CreateUser(ctx context.Context, user User) error {
	//insert into ddb
	return nil
}

func (r repo) GetUser(ctx context.Context, id string) (string, error) {
	//get user from db
	return "1", nil
}
