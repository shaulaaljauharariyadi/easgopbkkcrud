package models

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
	// CreatedAt time.Time `json:"createdate" gorm:"not null"`
	// UpdatedAt time.Time `json:"updatedate" gorm:"not null"`
}

func (Category) TableName() string {
	return "category"
}
