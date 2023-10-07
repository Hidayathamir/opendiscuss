package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	user := model.User{}

	q := ur.db.Table(model.USER_TABLE_NAME).
		Model(&model.User{}).
		Where(fmt.Sprintf("%s = ?", model.USER_USERNAME), username).
		First(&user)

	if q.Error != nil {
		return model.User{}, q.Error
	}

	return user, nil
}
