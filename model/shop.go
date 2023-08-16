package model

import (
	"gorm.io/gorm"
)

type Shop struct {
	ID          int64  // 主键
	ShopID      int64  // 店铺id
	ShopName    string // 店铺名称
	IsValid     int    // 是否有效
	AccountRole string // 账户角色

	AccessToken  string // access_token
	RefreshToken string // refresh_token
	gorm.Model
}

func (Shop) TableName() string {
	return "shop"
}
