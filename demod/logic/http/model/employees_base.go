package model

import (
	"demod/lib/db"
	"gorm.io/gorm"
)

// EmployeesBase struct is a row record of the employees_base table in the jz_ybs database
type EmployeesBase struct {
	BaseModel
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 5] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"` // 姓名
}

func (m EmployeesBase) Model() *gorm.DB {
	return db.Mysql((&m).Connection()).Model(&m)
}

func (*EmployeesBase) Connection() string {
	return "db"
}

func (m *EmployeesBase) TableName() string {
	return "db_name"
}
