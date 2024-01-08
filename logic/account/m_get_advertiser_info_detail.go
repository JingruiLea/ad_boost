package account

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"time"
)

func MGetAdvertiserInfoDetail(ctx context.Context, adIDs []int64) ([]*AdAccount, error) {
	params := map[string]interface{}{
		"advertiser_ids": adIDs,
	}
	var resp MGetAdInfoDetailRespData
	//adIDs[0] for access_token
	err := httpclient.NewClient().AdGet(ctx, adIDs[0], "https://ad.oceanengine.com/open_api/2/advertiser/info/", &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	return resp, nil
}

type MGetAdInfoDetailRespData []*AdAccount

type AdAccount struct {
	Address                 *string `json:"address"`
	Brand                   string  `json:"brand"`
	Company                 string  `json:"company"`
	CreateTime              string  `json:"create_time"`
	FirstIndustryName       string  `json:"first_industry_name"`
	ID                      int64   `json:"id"`
	Industry                string  `json:"industry"`
	LicenseCity             string  `json:"license_city"`
	LicenseNo               string  `json:"license_no"`
	LicenseProvince         string  `json:"license_province"`
	LicenseUrl              string  `json:"license_url"`
	Name                    string  `json:"name"`
	Note                    string  `json:"note"`
	PromotionArea           string  `json:"promotion_area"`
	PromotionCenterCity     string  `json:"promotion_center_city"`
	PromotionCenterProvince string  `json:"promotion_center_province"`
	Reason                  string  `json:"reason"`
	Role                    string  `json:"role"`
	SecondIndustryName      string  `json:"second_industry_name"`
	Status                  string  `json:"status"`
}

func (a *AdAccount) ToModel() *model.Advertiser {
	ret := &model.Advertiser{
		AdvertiserID:            a.ID,
		Name:                    a.Name,
		Company:                 a.Company,
		Brand:                   a.Brand,
		FirstIndustryName:       a.FirstIndustryName,
		Industry:                a.Industry,
		LicenseCity:             a.LicenseCity,
		LicenseNo:               a.LicenseNo,
		LicenseProvince:         a.LicenseProvince,
		LicenseURL:              a.LicenseUrl,
		Note:                    a.Note,
		PromotionArea:           a.PromotionArea,
		PromotionCenterCity:     a.PromotionCenterCity,
		PromotionCenterProvince: a.PromotionCenterProvince,
		Reason:                  a.Reason,
		Role:                    a.Role,
		SecondIndustryName:      a.SecondIndustryName,
		Status:                  a.Status,
	}
	createTime := parseTime(a.CreateTime)
	if !createTime.IsZero() {
		ret.CreateTime = createTime
	}
	if a.Address != nil {
		ret.Address = *a.Address
	}
	return ret
}

func parseTime(timeStr string) time.Time {
	if len(timeStr) == 0 {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
