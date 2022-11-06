package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product"`
	UserID    int                  `json:"user_id" gorm:"type:varchar(25)"`
	User      UsersProfileResponse `json:"user"`
	Price     int                  `json:"price" gorm:"type:int"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	Qty       int                  `json:"qty"`
	Date      string               `json:"date"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}
