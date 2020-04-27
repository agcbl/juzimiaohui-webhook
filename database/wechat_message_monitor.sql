CREATE TABLE `wechat_message_monitor` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `wxid` varchar(1024) NOT NULL,
  `wechat_name` varchar(1024) DEFAULT NULL,
  `room_name` varchar(1024) DEFAULT NULL,
  `content` varchar(1024) NOT NULL,
  `msg_type` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `room_id` varchar(200) DEFAULT '',
  `message_id` varchar(32) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_room` (`room_name`(191)),
  KEY `idx_wechat_name` (`wechat_name`(191))
) ENGINE=InnoDB AUTO_INCREMENT=1376 DEFAULT CHARSET=utf8mb4


CREATE TABLE `wechat_room` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `room_id` VARCHAR(191) NOT NULL ,
    `room_name` VARCHAR(1024) DEFAULT "",
    `room_member_number` INT(11) DEFAULT 0,
    `open_monitor` TINYINT(1) DEFAULT 1,
    PRIMARY KEY (`id`),
    UNIQUE (`room_id`),
    KEY `idx_room_name` (`room_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `wechat_user_info` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `wechat_id` VARCHAR(1024) DEFAULT "",
    `wxid` VARCHAR(1024) NOT NULL,
    `wechat_name` VARCHAR(1024) DEFAULT "",
    `gender` TINYINT(1) DEFAULT 0,
    `city` VARCHAR(200) DEFAULT "",
    `province` VARCHAR(200) DEFAULT "",
    `avatar_url` VARCHAR(1024) DEFAULT "",
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `last_active_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    PRIMARY KEY (`id`),
    KEY `idx_wxid` (`wxid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wechat_room_member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `wxid` varchar(1024) NOT NULL,
  `wechat_name` varchar(1024) DEFAULT '',
  `room_name` varchar(1024) DEFAULT '',
  `room_id` varchar(200) NOT NULL,
  `room_alias` varchar(1024) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_wechat_name` (`wechat_name`(191)),
  KEY `idx_room_id` (`room_id`(191))
) ENGINE=InnoDB AUTO_INCREMENT=36194 DEFAULT CHARSET=utf8mb4

CREATE TABLE `keyword` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `word` VARCHAR(50) NOT NULL,
    `is_opened` TINYINT(1) DEFAULT 1,
    PRIMARY KEY (`id`),
    UNIQUE `idx_word` (`word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `feishu_bot` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `token` VARCHAR(100) NOT NULL,
    `expire` BIGINT(11),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wechat_whitelist` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `wxid` VARCHAR(1024) DEFAULT "",
    PRIMARY KEY (`id`),
    KEY `idx_wxid` (`wxid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
