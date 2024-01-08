package bo

// DynamicWord represents a dynamic word object in the creative title information.
type DynamicWord struct {
	WordID      int    `json:"word_id"`      // 动态词包ID
	DictName    string `json:"dict_name"`    // 创意词包名称
	DefaultWord string `json:"default_word"` // 创意词包默认词
}

// ProgrammaticCreativeTitle represents a single creative title information.
type ProgrammaticCreativeTitle struct {
	Title           string        `json:"title"`             // 创意标题
	TitleType       string        `json:"title_type"`        // 标题类型
	AwemeCarouselID int           `json:"aweme_carousel_id"` // 抖音图文id
	DynamicWords    []DynamicWord `json:"dynamic_words"`     // 动态词包对象列表
}
