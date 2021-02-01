package service

import (
	"cicd/internal/model"
	"cicd/pkg/app"
)

/*
由于本项目并不复杂，所以直接把Request结构体放在了service层中以便使用。
若后续业务不断增长，程序越来越复杂，service层也变得冗杂，则可以考虑抽离一层
接口校验层，以便解耦逻辑
*/
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

//另外，我们还在service层中进行了一些简单的逻辑封装。
//在应用分层中，service层主要是对业务逻辑进行封装，
//如果有一些业务聚合和处理可以在该层进行编码，则可以较好地隔离上下两层的逻辑。
func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreadTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
