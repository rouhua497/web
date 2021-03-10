package dao

import "cicd/internal/model"



func (d *Dao) GetAuth(Appkey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    Appkey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
