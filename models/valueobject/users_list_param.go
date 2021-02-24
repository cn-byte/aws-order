package valueobject

/**
* @Description: 培训配置
**/
type UsersListParamVo struct {
	// 页码
	Page int32
	// 页的大小
	Size int32
	// 排序 asc desc
	CreateTimeOrder string
	// 状态
	Status []int32
}
