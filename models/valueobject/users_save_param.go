package valueobject

import "time"

/**
* @Description: 用户
**/
type UsersSaveParamVo struct {
	Id         int32      //主键
	UserName   string     //用户名
	Password   string     //密码
	Status     int32      //用户状态
	CreateTime *time.Time //创建时间
	UpdateTime *time.Time //修改时间
	CreateBy   string     //创建人
	UpdateBy   string     //修改人

}
