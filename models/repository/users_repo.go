package repository

import (
	"aws-order/database"
	"aws-order/models/dataobject"
	"aws-order/models/entity"
	"aws-order/models/valueobject"
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	// 顺序
	ASC string = "asc"
	// 倒序
	DESC string = "desc"
)

type UsersRepo struct {
}

// SelectByPage 分页查询
func (s *UsersRepo) SelectByPage(ctx context.Context, vParamVo valueobject.UsersListParamVo, page int32,
	limit int32) ([]entity.Users, int64, error) {
	var totalSize int64
	var res []entity.Users
	sqlBuild := ` SELECT u.*  FROM t_users AS u WHERE 1 = 1 `
	if len(vParamVo.Status) > 0 {
		sqlBuild += " AND u.status IN @Status"
	}
	sqlCount := fmt.Sprintf(" SELECT COUNT(*) FROM (%s) _t ", sqlBuild)
	database.TestDB.Raw(sqlCount, vParamVo).Count(&totalSize)
	var OrderFile []string
	if vParamVo.CreateTimeOrder == "" {
		OrderFile = append(OrderFile, "u.create_time desc ")
	}
	if vParamVo.CreateTimeOrder == ASC {
		OrderFile = append(OrderFile, " u.create_time asc ")
	}
	if vParamVo.CreateTimeOrder == DESC {
		OrderFile = append(OrderFile, " u.create_time desc ")
	}
	if len(OrderFile) > 0 {
		join := strings.Join(OrderFile, ",")
		sqlBuild += " Order By " + join
	}
	offset := (page - 1) * limit
	sqlBuild += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)
	database.TestDB.Debug().Raw(sqlBuild, vParamVo).Find(&res)
	return res, totalSize, nil
}

// Find 查询单个
func (s *UsersRepo) Find(ctx context.Context, users entity.Users) (*entity.Users, error) {
	var usersDO dataobject.UsersDO
	db := database.TestDB.Where(&users).First(&usersDO)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	var record entity.Users
	_ = gconv.Struct(usersDO, &record)
	return &record, db.Error
}

// Create 创建
func (s *UsersRepo) Create(ctx context.Context, users *entity.Users) error {
	now := time.Now()
	users.CreateTime = &now
	users.UpdateTime = &now
	var usersDO dataobject.UsersDO
	_ = gconv.Struct(users, &usersDO)
	err := database.TestDB.Create(&usersDO).Error
	if err != nil {
		return err
	}
	return nil
}

// Update 修改
func (s *UsersRepo) Update(ctx context.Context, users entity.Users) error {
	now := time.Now()
	users.UpdateTime = &now
	var usersDO dataobject.UsersDO
	_ = gconv.Struct(users, &usersDO)
	err := database.TestDB.Updates(&usersDO).Error
	if err != nil {
		return err
	}
	return nil
}
