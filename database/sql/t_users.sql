/*
Navicat MySQL Data Transfer

Source Server         : 本机
Source Server Version : 100417
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 100417
File Encoding         : 65001

Date: 2021-02-24 15:08:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_users
-- ----------------------------
DROP TABLE IF EXISTS `t_users`;
CREATE TABLE `t_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(32) NOT NULL,
  `status` int(255) NOT NULL DEFAULT 0,
  `create_by` varchar(50) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(50) DEFAULT '',
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

-- ----------------------------
-- Records of t_users
-- ----------------------------
