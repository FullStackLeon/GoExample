// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDep = "deps"

// Dep mapped from table <deps>
type Dep struct {
	ID   int32  `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Num  int32  `gorm:"column:num;not null" json:"num"`
}

// TableName Dep's table name
func (*Dep) TableName() string {
	return TableNameDep
}
