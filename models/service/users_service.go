package service

import (
	"aws-order/models/entity"
	"aws-order/models/enums"
	"aws-order/models/repository"
	"aws-order/models/valueobject"
	"aws-order/util"
	"context"
	"github.com/gogf/gf/util/gconv"
)

type UsersService struct {
	usersRepo repository.UsersRepo
}

// QueryList 查询列表
func (s *UsersService) QueryListPage(ctx context.Context, vParamVo valueobject.UsersListParamVo) ([]entity.Users, int64, error) {
	return s.usersRepo.SelectByPage(ctx, vParamVo, vParamVo.Page, vParamVo.Size)
}

// QueryDetail 查询详情
func (s *UsersService) QueryDetail(ctx context.Context, users entity.Users) (*entity.Users, error) {
	res, err := s.usersRepo.Find(ctx, users)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, util.NewError(404, "id 不存在")
	}
	return res, nil
}

//  Save 创建/修改
func (s *UsersService) Save(ctx context.Context, usersSaveParamVo valueobject.UsersSaveParamVo) error {
	var err error
	var users entity.Users
	_ = gconv.Struct(usersSaveParamVo, &users)
	if usersSaveParamVo.Id > 0 {
		find, _ := s.usersRepo.Find(ctx, entity.Users{Id: usersSaveParamVo.Id})
		if find == nil {
			return nil
		}
		err = s.usersRepo.Update(ctx, users)
	} else {
		users.Status = int32(enums.UsersStatusNormal)
		err = s.usersRepo.Create(ctx, &users)
	}
	if err != nil {
		return err
	}
	return nil
}

// ExistId 检查ID
func (s *UsersService) ExistId(ctx context.Context, id int32) error {
	res, err := s.usersRepo.Find(ctx, entity.Users{Id: id})
	if err != nil {
		return err
	}
	if res == nil {
		return util.NewError(404, "id 不存在")
	}
	return nil
}

// Disable  禁用
func (s *UsersService) Disable(ctx context.Context, users entity.Users) error {
	err := s.ExistId(ctx, users.Id)
	if err != nil {
		return err
	}
	err = s.usersRepo.Update(ctx, entity.Users{
		Id:     users.Id,
		Status: int32(enums.UsersStatusDisable),
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (s *UsersService) Delete(ctx context.Context, users entity.Users) error {
	err := s.ExistId(ctx, users.Id)
	if err != nil {
		return err
	}
	err = s.usersRepo.Update(ctx, entity.Users{
		Id:     users.Id,
		Status: int32(enums.UsersStatusDelete),
	})
	if err != nil {
		return err
	}
	return nil
}
