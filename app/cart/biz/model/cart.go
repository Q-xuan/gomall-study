package model

import (
	"context"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null; index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null";`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart *Cart) error {
	return db.Create(cart).Error
	
}
