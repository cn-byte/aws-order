package enums

/**
* @Description:
* @Author: LiuXianSong
* @Date: 2020/02/01 8:37 下午
**/

// 用户状态类型
type UsersStatusEnum int32

const (
	// 正常
	UsersStatusNormal UsersStatusEnum = 1
	// 禁用
	UsersStatusDisable UsersStatusEnum = 2
	// 删除
	UsersStatusDelete UsersStatusEnum = 3
)
