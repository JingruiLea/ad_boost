
CREATE TABLE delivery_setting (
      id BIGINT NOT NULL PRIMARY KEY COMMENT '主键, 广告表ID',
      ad_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告ID',
      delivery_setting_budget DECIMAL(10,2)  NOT NULL DEFAULT 0 COMMENT '预算',
      delivery_setting_budget_mode VARCHAR(255) NOT NULL DEFAULT '' COMMENT '预算模式',
      delivery_setting_deep_bid_type VARCHAR(255) DEFAULT NULL COMMENT '深度出价类型',
      delivery_setting_deep_external_action VARCHAR(255) DEFAULT NULL COMMENT '深度优化指标',
      delivery_setting_end_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '结束时间',
      delivery_setting_external_action VARCHAR(255) NOT NULL DEFAULT '' COMMENT '优化指标',
      delivery_setting_product_new_open BOOLEAN NOT NULL DEFAULT false COMMENT '新产品开放',
      delivery_setting_roi_goal DECIMAL(10,2)  DEFAULT NULL COMMENT 'ROI目标',
      delivery_setting_smart_bid_type VARCHAR(255) NOT NULL DEFAULT '' COMMENT '智能出价类型',
      delivery_setting_start_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '开始时间',
      delivery_setting_cpa_bid DECIMAL(10,2)  DEFAULT NULL COMMENT 'CPA出价',
      delivery_setting_advertiser_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告主ID',

      created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
      updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
      deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',

      UNIQUE KEY (ad_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '投放设置';

