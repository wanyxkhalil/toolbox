CREATE TABLE `user`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL COMMENT '用户名',
    `valid`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有效：0_无效，1_有效',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=17318 DEFAULT CHARSET=utf8mb4 COMMENT='用户';