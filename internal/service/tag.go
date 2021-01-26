package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//tag list列表
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//新增tag
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//修改tag
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"` //binding：入参校验的规则内容
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

//删除tag
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
