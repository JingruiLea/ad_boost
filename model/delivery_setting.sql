
CREATE TABLE delivery_setting (
      id BIGINT NOT NULL PRIMARY KEY COMMENT '主键, 广告表ID',
      ad_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告ID',
      budget DECIMAL(10,2)  NOT NULL DEFAULT 0 COMMENT '预算',
      budget_mode VARCHAR(255) NOT NULL DEFAULT '' COMMENT '预算模式',
      deep_bid_type VARCHAR(255) DEFAULT NULL COMMENT '深度出价类型',
      deep_external_action VARCHAR(255) DEFAULT NULL COMMENT '深度优化指标',
      end_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '结束时间',
      external_action VARCHAR(255) NOT NULL DEFAULT '' COMMENT '优化指标',
      product_new_open BOOLEAN NOT NULL DEFAULT false COMMENT '新产品开放',
      roi_goal DECIMAL(10,2)  DEFAULT NULL COMMENT 'ROI目标',
      smart_bid_type VARCHAR(255) NOT NULL DEFAULT '' COMMENT '智能出价类型',
      start_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '开始时间',
      cpa_bid DECIMAL(10,2)  DEFAULT NULL COMMENT 'CPA出价',

      created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
      updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
      deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',

      UNIQUE KEY (ad_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '投放设置';