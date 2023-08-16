package model

import (
	"gorm.io/gorm"
	"time"
)

type Advertiser struct {
	AdvertiserID            int64     `gorm:"unique;notNull" json:"advertiser_id"`
	Name                    string    `json:"name"`
	Company                 string    `json:"company"`
	Address                 string    `json:"address"`
	Brand                   string    `json:"brand"`
	CreateTime              time.Time `json:"create_time"`
	FirstIndustryName       string    `json:"first_industry_name"`
	Industry                string    `json:"industry"`
	LicenseCity             string    `json:"license_city"`
	LicenseNo               string    `json:"license_no"`
	LicenseProvince         string    `json:"license_province"`
	LicenseURL              string    `json:"license_url"`
	Note                    string    `json:"note"`
	PromotionArea           string    `json:"promotion_area"`
	PromotionCenterCity     string    `json:"promotion_center_city"`
	PromotionCenterProvince string    `json:"promotion_center_province"`
	Reason                  string    `json:"reason"`
	Role                    string    `json:"role"`
	SecondIndustryName      string    `json:"second_industry_name"`
	Status                  string    `json:"status"`

	gorm.Model
}

func (a *Advertiser) TableName() string {
	return "advertiser"
}
