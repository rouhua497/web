package model

import (
	"cicd/global"
	"cicd/pkg/setting"
	"fmt"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"` //只表示正数
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

//初始化gorm
func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParesTime,
		))

	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, err

}

/*
· 通过调用scope.FieldByName方法，获取当前是否包含所需的字段。
· 通过判断Field.IsBlank的值，可以得知该字段的值是否为空。
· 若为空，则调用Field.Set方法给该字段设置值。入参类型为interface{}，即内部是通过反射进行一系列操作赋值的
*/

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

/*
· 通过调用 scope.Get（″gorm：update_column″） 来获取当前设置的标识gorm：update_column的字段属性。
· 若不存在，即没有自定义设置 update_column，则在更新回调内设置默认字段ModifiedOn的值为当前的时间戳。
*/
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

/*
·通过调用scope.Get（″gorm：delete_option″）来获取当前设置的标识gorm：delete_option的字段属性。

·判断是否存在DeletedOn和IsDel字段。若存在，则执行UPDATE操作进行软删除（修改DeletedOn和IsDel的值），
 否则执行DELETE操作进行硬删除。

·调用scope.QuotedTableName方法获取当前引用的表名，并调用一系列方法对SQL语句的组成部分进行处理和转移。
 在完成一些所需参数设置后，调用scope.CombinedConditionSql方法完成SQL语句的组装
*/
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
