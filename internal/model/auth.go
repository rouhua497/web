package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key=? and app_secret = ? and is_del = ? ", a.AppKey, a.AppSecret, 0)
	if err := db.First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}
