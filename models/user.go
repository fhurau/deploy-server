package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"type: varchar(50)"`
	Phone     string    `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Location  string    `json:"location" form:"location" gorm:"type: varchar(255)"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Role      string    `json:"role" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Role     string `json:"role"`
	Gender   string `json:"gender"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
