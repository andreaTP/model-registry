// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameTypeProperty = "TypeProperty"

// TypeProperty mapped from table <TypeProperty>
type TypeProperty struct {
	TypeID   int64  `gorm:"column:type_id;primaryKey" json:"-"`
	Name     string `gorm:"column:name;primaryKey;not null" json:"-"`
	DataType *int32 `gorm:"column:data_type" json:"-"`
}

// TableName TypeProperty's table name
func (*TypeProperty) TableName() string {
	return TableNameTypeProperty
}