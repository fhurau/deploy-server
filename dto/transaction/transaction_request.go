package transaction

type TransactionRequest struct {
	UserID    int    `json:"user_id" gorm:"type: int"`
	ProductID int    `json:"product_id" form:"product_id" gorm:"type: int"`
	Price     int    `json:"price" form:"price" gorm:"type: int"`
	Status    string `json:"status" form:"status" gorm:"type: varchar(255)"`
}
type UpdateTransactionRequest struct {
	UserID    int    `json:"user_id" gorm:"type: int"`
	ProductID int    `json:"product_id" form:"product_id" gorm:"type: int"`
	Price     int    `json:"price" form:"price" gorm:"type: int"`
	Status    string `json:"status" form:"status" gorm:"type: varchar(255)"`
}
