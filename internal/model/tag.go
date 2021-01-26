package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

//获取tag数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db.Where("name=? ", t.Name)
	}
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		//Offset指定开始返回记录前要跳过的记录数。
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if len(t.Name) > 0 {
		db.Where("name =  ?", t.Name)
	}
	db = db.Where("state=?", t.State)
	if err = db.Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, err
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Model：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，
//格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的TableName方法，
//用于指定其对应返回的表名
func (t Tag) Update(db *gorm.DB) error {
	db = db.Model(&Tag{}).Where("id =? and is_del=?", t.ID, 0)
	return db.Update(t).Error
}
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.Model.ID, 0).Delete(&t).Error
}
