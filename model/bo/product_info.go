package bo

type ProductInfo struct {
	ID                  int     `json:"id"`                    // 商品id
	Name                string  `json:"name"`                  // 商品名称
	DiscountPrice       float64 `json:"discount_price"`        // 售价，已废弃
	Img                 string  `json:"img"`                   // 商品主图
	MarketPrice         float64 `json:"market_price"`          // 原价，单位为元
	DiscountLowerPrice  float64 `json:"discount_lower_price"`  // 折扣价区间最小值，单位为元
	DiscountHigherPrice float64 `json:"discount_higher_price"` // 折扣价区间最大值，单位为元
	ChannelID           int     `json:"channel_id"`            // 渠道ID
	ChannelType         string  `json:"channel_type"`          // 渠道类型
}
