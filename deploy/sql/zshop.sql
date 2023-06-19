DROP DATABASE IF EXISTS `zshop`;
CREATE DATABASE zshop charset=utf8mb4;
USE zshop;
/*用户*/
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` char(100) NOT NULL COMMENT '用户名',
    `mobile` char(100) NOT NULL COMMENT '手机号',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*商品*/
CREATE TABLE `goods` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` char(100) NOT NULL COMMENT '商品名',
    `price` float NOT NULL COMMENT '价格',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*库存*/
CREATE TABLE `inventory` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `goods` bigint NOT NULL COMMENT '商品id',
    `stocks` bigint NOT NULL COMMENT '库存数量',
    `freeze` bigint NOT NULL DEFAULT '0' COMMENT '冻结数量',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_goods` (`goods`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*订单1*/
CREATE TABLE `order_info` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user` bigint NOT NULL COMMENT '用户id',
    `order_sn` char(100) NOT NULL COMMENT '订单号',
    `statue` char(100) NOT NULL DEFAULT '交易创建' COMMENT '订单状态',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_order_sn` (`order_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*订单2*/
CREATE TABLE `order_goods` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `order_sn` char(100) NOT NULL COMMENT '订单号',
    `goods` bigint NOT NULL COMMENT '商品id',
    `price` float NOT NULL COMMENT '价格',
    `nums` bigint NOT NULL COMMENT '商品数量',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/**/
INSERT INTO zshop.goods (id, name, price) VALUES (1, '苹果', 10);
INSERT INTO zshop.goods (id, name, price) VALUES (2, '香蕉', 11);
INSERT INTO zshop.goods (id, name, price) VALUES (3, '西瓜', 12);

INSERT INTO zshop.inventory (id, goods, stocks, freeze) VALUES (1, 1, 1000, 0);
INSERT INTO zshop.inventory (id, goods, stocks, freeze) VALUES (2, 2, 1000, 0);
INSERT INTO zshop.inventory (id, goods, stocks, freeze) VALUES (3, 3, 1000, 0);

INSERT INTO zshop.user (id, name, mobile) VALUES (1, 'myboran', '10086');
