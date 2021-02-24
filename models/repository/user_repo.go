package repository

import (
	"aws-order/models/entity"
	"context"
)

type UserRepo struct {
}

// Find 查询
func (s *UserRepo) Find(ctx context.Context, id uint32) (res []entity.Users, err error) {

	return
}
