// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameChild = "child"

// Child mapped from table <child>
type Child struct {
	ID   uint32 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

// TableName Child's table name
func (*Child) TableName() string {
	return TableNameChild
}
