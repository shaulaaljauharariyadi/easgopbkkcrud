package models

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
}

func (Category) TableName() string {
	return "category"
}
