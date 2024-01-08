package boost

// AdBoostParams 广告优化的超参数
type AdBoostParams struct {
	MinROI  float64 `json:"min_roi"`  // 最小ROI
	MaxROI  float64 `json:"max_roi"`  // 最大ROI
	InitROI float64 `json:"init_roi"` // 初始ROI

	MinBid  float64 `json:"min_bid"`  // 最小出价
	MaxBid  float64 `json:"max_bid"`  // 最大出价
	InitBid float64 `json:"init_bid"` // 初始出价

	MinBudget float64 `json:"min_budget"` // 最小预算
	MaxBudget float64 `json:"max_budget"` // 最大预算
}
