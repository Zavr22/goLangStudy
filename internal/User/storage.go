package User

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user *User) error
	FindAll(ctx context.Context) (u []User, err error)
	FindOne(ctx context.Context, dataname string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
