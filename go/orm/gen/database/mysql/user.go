package mysql

import (
	"context"

	"github.com/tusmasoma/samples/go/orm/gen/database"
	"github.com/tusmasoma/samples/go/orm/gen/database/mysql/model"
	"github.com/tusmasoma/samples/go/orm/gen/database/mysql/query"
	"github.com/tusmasoma/samples/go/orm/gen/entity"
	"gorm.io/gorm"
)

type user struct {
	query *query.Query
}

func NewUser(db *gorm.DB) database.User {
	return &user{
		query: query.Use(db),
	}
}

func (u *user) Get(ctx context.Context, id string) (*entity.User, error) {
	userModel, err := u.query.User.WithContext(ctx).Where(u.query.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	user, err := entity.NewUser(userModel.ID, userModel.Name, userModel.Email, userModel.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) Create(ctx context.Context, user *entity.User) error {
	userModel := &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.query.User.WithContext(ctx).Create(userModel); err != nil {
		return err
	}
	return nil
}
func (u *user) Update(ctx context.Context, user *entity.User) error {
	userModel := &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.query.User.WithContext(ctx).Save(userModel); err != nil {
		return err
	}
	return nil
}
func (u *user) Delete(ctx context.Context, id string) error {
	if _, err := u.query.User.WithContext(ctx).Where(u.query.User.ID.Eq(id)).Delete(); err != nil {
		return err
	}
	return nil
}
