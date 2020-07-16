/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.8.105
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : 192.168.8.105:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 16/07/2020 07:57:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `id_card` varchar(255) DEFAULT NULL COMMENT '身份证',
  `eid` varchar(0) DEFAULT NULL COMMENT '国安三所id地址',
  `chain_address` varchar(255) DEFAULT NULL COMMENT '伪的区块链地址',
  `nt_chain_address` varchar(255) DEFAULT NULL COMMENT '全民链地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_phone` (`phone`) COMMENT '手机号码唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
