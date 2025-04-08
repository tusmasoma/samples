package database

import (
	"context"

	"github.com/tusmasoma/samples/go/orm/gen/entity"
)

type Database struct {
	User User
}

type User interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
