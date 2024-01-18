package ttypes

type BaseResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalNumber int `json:"total_number"`
	TotalPage   int `json:"total_page"`
}

type BoolInt int

const (
	BoolIntFalse BoolInt = 0
	BoolIntTrue  BoolInt = 1
)
