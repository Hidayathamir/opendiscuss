package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (ur *UserRepository) RegisterUser(ctx context.Context, user model.User) (int, error) {
	q := ur.db.Table(model.USER_TABLE_NAME).Create(&user)
	if q.Error != nil {
		return 0, q.Error
	}
	return user.ID, nil
}
