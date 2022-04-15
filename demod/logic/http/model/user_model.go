/**
** @创建时间 : 2022/1/6 10:52
** @作者 : fzy
 */
package model

import (
	"demod/lib/db"
	"gorm.io/gorm"
)

type UserModel struct {
	BaseModel
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`

	Username string `gorm:"column:username;type:varchar;size:64;" json:"username"` // 姓名
}

func (m UserModel) Model() *gorm.DB {
	return db.Mysql((&m).Connection()).Model(&m)
}

func (*UserModel) Connection() string {
	return "db"
}

func (m *UserModel) TableName() string {
	return "user"
}