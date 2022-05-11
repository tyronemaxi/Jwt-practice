CREATE TABLE `user`
(
    `pk`              bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`         varchar(32)  NOT NULL COMMENT '用户唯一 ID',
    `domain_id`       varchar(32)  NOT NULL COMMENT '云账户 ID',
    `username`        varchar(128) NOT NULL COMMENT '用户名',
    `password`        varchar(256) NOT NULL COMMENT '密码',
    `display_name`    varchar(256) NOT NULL DEFAULT '' COMMENT '别名',
    `source_type`     varchar(64)  NOT NULL DEFAULT '' COMMENT '来源：create/import',
    `mobile_phone`    varchar(256) NOT NULL DEFAULT '' COMMENT '用户手机号',
    `mobile_verified` tinyint(1)   NOT NULL DEFAULT 0 COMMENT '用户手机号是否验证：0/1',
    `email`           varchar(256) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `email_verified`  tinyint(1)   NOT NULL DEFAULT 0 COMMENT '邮箱是否验证：0/1',
    `extra`           text         NOT NULL COMMENT '扩展字段',
    `create_time`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time`    datetime              DEFAULT NULL COMMENT '删除时间',
    `is_deleted`      tinyint(1)   NOT NULL DEFAULT 0 COMMENT '是否删除：0-未删除，1-删除',
    PRIMARY KEY (`pk`),
    UNIQUE KEY `user_id` (`user_id`, `mobile_phone`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8 COMMENT ='用户信息表'