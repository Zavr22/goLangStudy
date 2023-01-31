package db

import (
	"context"
	"fmt"
	"github.com/Zavr22/goLangStudy/cmd/pkg/postgresql"
	"github.com/Zavr22/goLangStudy/internal/User"
	"github.com/jackc/pgconn"
	"log"
	"strings"
)

type Repository struct {
	client postgresql.Client
	logger *log.Logger
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", ""), "\n", "")
}

func (r *Repository) CreateUser(ctx context.Context, user User.User) (string, error) {
	query := `
				INSERT INTO "user"
				( name, surname, password, role)
				VALUES ($1, $2, $3, $4)
				RETURNING id
	`
	r.logger.Println(fmt.Sprintf("Sql query: %s", formatQuery(query)))
	if err := r.client.QueryRow(ctx, query, user.Name, user.Surname, user.Password, user.Role).Scan(&user.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newError := fmt.Sprintf(" Sql Error: %s, Detail: %s, Where : %s, Code: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			fmt.Println(newError)
			return "", nil
		}
		return "", err
	}
	return "", nil
}

func (r *Repository) FindAll(ctx context.Context) (u []User.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) FindOne(ctx context.Context) (User.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(ctx context.Context, user User.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *log.Logger) User.Storage {
	return &Repository{
		client: client,
		logger: logger,
	}
}
