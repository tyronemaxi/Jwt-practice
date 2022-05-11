# 用户注册
前端提交的信息
- 用户名（必填）
- 用户密码 （必填）
- 别名（非必填）
- 手机号（必填）
- 邮箱（非必须填）

# api 设计
`user/register`
{
"username": "tyrone",
"password": "teedfadfa",
"mobile_phone": "13041479216",
"email": "tyrone@gmail.com"
}
## 正则
```go
const (
	UserNameRule    = "^[\u4e00-\u9fa5a-zA-Z0-9]{5,20}$"
	PasswordRule    = "^([A-Z]|[a-z]|[0-9]|[-=[;,./~!@#$%^*()_+}{:?]){6,20}$"
	MobilePhoneRule = "^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"
	EmailRule       = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
)
```
## 用户表设计
```mysql
CREATE TABLE `user`
(
    `pk`              bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`         varchar(32)  NOT NULL COMMENT '用户唯一 ID',
    `domain_id`       varchar(32)  NOT NULL COMMENT '云账户 ID',
    `project_id`      varchar(32)  NOT NULL COMMENT '项目 ID',
    `username`        varchar(128) NOT NULL COMMENT '用户名',
    `password`        varchar(256) NOT NULL COMMENT '密码',
    `display_name`    varchar(256) NOT NULL DEFAULT '' COMMENT '别名',
    `source_type`     varchar(64)  NOT NULL DEFAULT '' COMMENT '来源：create/import',
    `mobile_phone`    varchar(256) NOT NULL DEFAULT '' COMMENT '用户手机号',
    `mobile_verified` tinyint(1)   NOT NULL DEFAULT 0 COMMENT '用户手机号是否验证：0/1',
    `email`           varchar(256) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `email_verified`  tinyint(1)   NOT NULL DEFAULT 0 COMMENT '邮箱是否验证：0/1',
    `extra`           text         NOT NULL COMMENT '扩展字段',
    `create_time`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by`       varchar(128) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_time`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_by`       varchar(128) NOT NULL DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`pk`),
    UNIQUE KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8 COMMENT ='用户信息表'
```

