CREATE TABLE ad (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    ad_create_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '广告创建时间',
    ad_id BIGINT NOT NULL COMMENT '广告ID',
    ad_modify_time VARCHAR(255) NOT NULL DEFAULT '' COMMENT '广告修改时间',
    campaign_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告组ID',
    campaign_scene VARCHAR(255) NOT NULL DEFAULT '' COMMENT '广告组场景',
    lab_ad_type VARCHAR(255) NOT NULL DEFAULT '' COMMENT '实验广告类型',
    marketing_goal VARCHAR(255) NOT NULL DEFAULT '' COMMENT '市场营销目标',
    marketing_scene VARCHAR(255) NOT NULL DEFAULT '' COMMENT '市场营销场景',
    name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '名称',
    opt_status VARCHAR(255) NOT NULL DEFAULT '' COMMENT '优化状态',
    status VARCHAR(255) NOT NULL DEFAULT '' COMMENT '状态',
    advertiser_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告主ID',

    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',

    UNIQUE KEY (ad_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '广告信息';

# add column advertiser_id
ALTER TABLE ad ADD COLUMN advertiser_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告主ID';
ALTER TABLE ad ADD INDEX idx_advertiser_id (advertiser_id);
