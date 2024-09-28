package model

type Parent struct {
	ID   uint32 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null;" json:"name"`

	Child Child `gorm:"foreignKey:ID" json:"child"`
}

func (p *Parent) TableName() string {
	return "parents"
}
