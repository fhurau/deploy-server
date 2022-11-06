package productdto

import "time"

type ProductResponse struct {
	ID       int               `json:"id"`
	Name     string            `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc     string            `json:"desc" gorm:"type:text" form:"desc"`
	Price    int               `json:"price" form:"price" gorm:"type: int"`
	Image    string            `json:"image" form:"image"`
	Qty      int               `json:"qty" form:"qty" gorm:"type: int"`
	UserID   int               `json:"user_id" gorm:"type: int"`
	Category []CategoryRespond `json:"category" form:"category" gorm:"type: int"`
	// CategoryID []int             `json:"-"`
}

type CategoryRespond struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
