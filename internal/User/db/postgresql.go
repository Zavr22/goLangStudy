package User

import (
	"context"
	"fmt"
	"github.com/Zavr22/goLangStudy/cmd/pkg/postgresql"
	user "github.com/Zavr22/goLangStudy/internal/User"
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

func (r *Repository) CreateUser(ctx context.Context, user user.User) error {
	query := `
				INSERT INTO PUBLIC."user"
				( name, surname, password, role)
				VALUES ($1, $2, $3, $4)
				RETURNING id
	`
	//r.logger.Println(fmt.Sprintf("Sql query: %s", formatQuery(query)))
	if err := r.client.QueryRow(ctx, query, user.Name, user.Surname, user.Password, user.Role).Scan(&user.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newError := fmt.Sprintf(" Sql Error: %s, Detail: %s, Where : %s, Code: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			fmt.Println(newError)
			return nil
		}
		return err
	}
	return nil
}

func (r *Repository) FindAll(ctx context.Context) (u []user.User, err error) {
	query := `
		SELECT USER.id, USER.name, USER.surname FROM PUBLIC.USER;
		`
	rows, err := r.client.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	users := make([]user.User, 0)

	for rows.Next() {
		var usr user.User

		err = rows.Scan(
			&usr.Id,
			&usr.Name,
			&usr.Surname,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) FindOne(ctx context.Context) (user.User, error) {

	query := `SELECT USER.ID, USER.NAME, USER.SURNAME FROM PUBLIC.USER WHERE USER.name= $1`

	r.logger.Println(fmt.Sprintf("Sql query: %s", formatQuery(query)))

	var usr user.User
	err := r.client.QueryRow(ctx, query, usr.Id).Scan(&usr.Id, &usr.Name, &usr.Surname)

	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (r *Repository) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client) *Repository {
	return &Repository{
		client: client,
	}
}
