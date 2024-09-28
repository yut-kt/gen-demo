package model

type Child struct {
	ID   uint32 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null;" json:"name"`
}

func (c *Child) TableName() string {
	return "child"
}
