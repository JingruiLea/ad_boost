# 店铺管理
CREATE TABLE shop (
                      id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                      shop_id BIGINT NOT NULL DEFAULT 0 COMMENT '店铺id',
                      shop_name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '店铺名称',
                      is_valid TINYINT NOT NULL DEFAULT 0 COMMENT '是否有效',
                      account_role VARCHAR(255) NOT NULL DEFAULT '' COMMENT '账户角色',

                      access_token VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'access_token',
                      refresh_token VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'refresh_token',

                      created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
                      updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
                      deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',

                      UNIQUE KEY (shop_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '店铺管理';
