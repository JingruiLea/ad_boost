#广告账户管理
CREATE TABLE advertiser (
                            id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                            advertiser_id BIGINT NOT NULL DEFAULT 0 COMMENT '广告主id',
                            name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '广告主名称',
                            company VARCHAR(255) NOT NULL DEFAULT '' COMMENT '公司名称',
                            address VARCHAR(255) DEFAULT NULL COMMENT '地址',
                            brand VARCHAR(255) DEFAULT '' COMMENT '品牌',
                            create_time DATETIME DEFAULT NULL COMMENT '创建时间',
                            first_industry_name VARCHAR(255) DEFAULT '' COMMENT '一级行业名称',
                            industry VARCHAR(255) DEFAULT '' COMMENT '行业',
                            license_city VARCHAR(255) DEFAULT '' COMMENT '许可证城市',
                            license_no VARCHAR(255) DEFAULT '' COMMENT '许可证号',
                            license_province VARCHAR(255) DEFAULT '' COMMENT '许可证省份',
                            license_url VARCHAR(255) DEFAULT '' COMMENT '许可证URL',
                            note VARCHAR(255) DEFAULT '' COMMENT '备注',
                            promotion_area VARCHAR(255) DEFAULT '' COMMENT '推广区域',
                            promotion_center_city VARCHAR(255) DEFAULT '' COMMENT '推广中心城市',
                            promotion_center_province VARCHAR(255) DEFAULT '' COMMENT '推广中心省份',
                            reason VARCHAR(255) DEFAULT '' COMMENT '原因',
                            role VARCHAR(255) DEFAULT '' COMMENT '角色',
                            second_industry_name VARCHAR(255) DEFAULT '' COMMENT '二级行业名称',
                            status VARCHAR(255) DEFAULT '' COMMENT '状态',

                            created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
                            updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
                            deleted_at DATETIME(3) NULL DEFAULT NULL,

                            UNIQUE KEY (advertiser_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '广告账户管理';