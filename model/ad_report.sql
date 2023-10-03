CREATE TABLE ad_report_item (
    id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    ad_id bigint(20) NOT NULL DEFAULT '0' COMMENT '广告计划id',
    advertiser_id bigint(20) NOT NULL DEFAULT '0' COMMENT '广告主id',
    click_cnt int(11) NOT NULL DEFAULT '0' COMMENT '点击次数',
    convert_cnt int(11) NOT NULL DEFAULT '0' COMMENT '转化数',
    convert_cost DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '转化成本',
    convert_rate DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '转化率',
    cpm_platform DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '平均千次展示费用',
    ctr DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '点击率',
    dy_follow int(11) NOT NULL DEFAULT '0' COMMENT '新增粉丝数',
    pay_order_amount DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '直接成交金额',
    pay_order_count int(11) NOT NULL DEFAULT '0' COMMENT '直接成交订单数',
    prepay_and_pay_order_roi DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '直接支付roi',
    show_cnt int(11) NOT NULL DEFAULT '0' COMMENT '展示次数',
    stat_cost DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '消耗',

    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间'

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告报表';

ALTER TABLE ad_report_item ADD INDEX idx_ad_id(ad_id);
ALTER TABLE ad_report_item ADD INDEX idx_advertiser_id(advertiser_id);
ALTER TABLE ad_report_item ADD INDEX idx_created_at(created_at);
ALTER TABLE ad_report_item ADD INDEX idx_updated_at(updated_at);
ALTER TABLE ad_report_item ADD INDEX idx_deleted_at(deleted_at);

ALTER TABLE `ad_report_item` ADD COLUMN `cpa_bid` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '出价' AFTER `stat_cost`;
ALTER TABLE `ad_report_item` ADD COLUMN `roi_goal` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT 'roi' AFTER `cpa_bid`;