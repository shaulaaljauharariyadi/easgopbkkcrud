package models

type Product struct {
	Id          uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string   `json:"name" gorm:"not null"`
	Category_Id uint     `json:"-" gorm:"not null"`
	Category    Category `json:"category" gorm:"foreignkey:category_id;references:id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Stock       int      `json:"stock" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
}

func (Product) TableName() string {
	return "product"
}
