package mysql

import (
	"context"

	"github.com/tusmasoma/samples/go/orm/gen/entity"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func (u *user) Get(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (u *user) Create(ctx context.Context, user *entity.User) error {
	return nil
}
func (u *user) Update(ctx context.Context, user *entity.User) error {
	return nil
}
func (u *user) Delete(ctx context.Context, id string) error {
	return nil
}
