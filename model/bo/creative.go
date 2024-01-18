package bo

import "github.com/JingruiLea/ad_boost/model/ttypes"

type Creative struct {
	CreativeCreateTime string           `json:"creative_create_time"`
	CreativeId         int64            `json:"creative_id"`
	CreativeModifyTime string           `json:"creative_modify_time"`
	ImageMode          ttypes.ImageMode `json:"image_mode"`
}
