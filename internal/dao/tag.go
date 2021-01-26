package dao

import (
	"cicd/internal/model"
	"cicd/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d Dao) CreadTag(name string, state uint8, createdBy uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			CreatedOn: createdBy,
		},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}
