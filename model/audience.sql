CREATE TABLE `audience` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `audience_id` bigint NOT NULL COMMENT '人群包ID',
    `name` varchar(255) NOT NULL COMMENT '定向包名称',
    `config` json DEFAULT NULL COMMENT '配置',

    `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY `audience_id` (`audience_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='人群包';