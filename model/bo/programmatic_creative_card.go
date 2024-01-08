package bo

type ProgrammaticCreativeCard struct {
	PromotionCardID                      int      `json:"promotion_card_id"`                        // 推广卡片ID
	ComponentID                          int      `json:"component_id"`                             // 组件ID
	PromotionCardTitle                   string   `json:"promotion_card_title"`                     // 推广卡片标题
	PromotionCardSellingPoints           []string `json:"promotion_card_selling_points"`            // 推广卡片卖点列表
	PromotionCardImageID                 string   `json:"promotion_card_image_id"`                  // 推广卡片配图ID
	PromotionCardActionButton            string   `json:"promotion_card_action_button"`             // 推广卡片行动号召按钮文案
	PromotionCardButtonSmartOptimization int      `json:"promotion_card_button_smart_optimization"` // 智能优选行动号召按钮文案开关
}
