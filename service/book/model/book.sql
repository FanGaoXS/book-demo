CREATE TABLE `book` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '书名',
                        `type` varchar(255)  NOT NULL COMMENT '书的分类',
                        `count` int NOT NULL COMMENT '书的数量',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `number_unique` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;