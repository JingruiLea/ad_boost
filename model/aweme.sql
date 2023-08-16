CREATE TABLE aweme (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',

    aweme_avatar VARCHAR(255) NOT NULL COMMENT 'Aweme头像',
    aweme_has_live_permission BOOLEAN NOT NULL COMMENT 'Aweme是否有直播权限',
    aweme_has_uni_prom BOOLEAN NOT NULL COMMENT 'Aweme是否有Uni Prom',
    aweme_has_video_permission BOOLEAN NOT NULL COMMENT 'Aweme是否有视频权限',
    aweme_id BIGINT NOT NULL COMMENT 'Aweme ID',
    aweme_name VARCHAR(255) NOT NULL COMMENT 'Aweme名称',
    aweme_show_id VARCHAR(255) NOT NULL COMMENT 'Aweme展示ID',
    aweme_status VARCHAR(255) NOT NULL COMMENT 'Aweme状态',
    bind_type JSON NOT NULL COMMENT '绑定类型',

    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',

    UNIQUE KEY (aweme_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT 'Aweme表';