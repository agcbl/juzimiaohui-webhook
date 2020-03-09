# juzimiaohui-webhook
句子秒回 Webhook

## Config

```toml
[database]
name = "dbname"
user = "user"
password = "password"
host = "db host"

[lark]
path = "bot path"

[juzihudong]
endpoint = "api"
token = "token"

[keyword]
sync_timer = 5
```

## Run

### Init Database

```sql
CREATE TABLE `wechat_message_monitor` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `wechat_id` varchar(1024) NOT NULL,
  `wechat_name` varchar(1024) DEFAULT NULL,
  `room_name` varchar(1024) DEFAULT NULL,
  `content` varchar(1024) NOT NULL,
  `msg_type` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `room_id` varchar(200) DEFAULT '',
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
    `room_id` VARCHAR(191) NOT NULL,
    `wxid` VARCHAR(1024) NOT NULL,
    `wechat_name` VARCHAR(1024) DEFAULT "",
    `gender` TINYINT(1) DEFAULT 0,
    `city` VARCHAR(200) DEFAULT "",
    `province` VARCHAR(200) DEFAULT "",
    `avatar_url` VARCHAR(1024) DEFAULT "",
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `last_active_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    PRIMARY KEY (`id`),
    KEY `idx_wxid_room_id` (`wxid`, `room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wechat_keywords` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `word` VARCHAR(50) NOT NULL,
    `is_opened` TINYINT(1) DEFAULT 1,
    PRIMARY KEY (`id`),
    UNIQUE `idx_word` (`word`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### Init room

```shell script
go run scripts/room/main.go -token Token -config /path/to/webhook.toml
```

### Run webhook service

```shell script
docker run -p 8000:8000 -v /path/to/webhook.toml:/etc/webhook.toml fatelei/juzhimiaohui-webhook:1.1
```
