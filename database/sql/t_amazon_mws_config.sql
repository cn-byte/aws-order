/*
Navicat MySQL Data Transfer

Source Server         : 本机
Source Server Version : 100417
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 100417
File Encoding         : 65001

Date: 2021-02-24 15:40:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_amazon_mws_config
-- ----------------------------
DROP TABLE IF EXISTS `t_amazon_mws_config`;
CREATE TABLE `t_amazon_mws_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` int(11) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `developers` varchar(32) NOT NULL DEFAULT '' COMMENT '开发者ID（12位数字的标识符',
  `access_key` varchar(32) NOT NULL DEFAULT '' COMMENT 'AWS Access Key ID（20个字符的字母数字标识符',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密钥（40个字符的标识符）',
  `status` int(255) NOT NULL DEFAULT 1 COMMENT '状态',
  `desc` varchar(100) DEFAULT '' COMMENT '描述信息',
  `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(50) DEFAULT '' COMMENT '修改人',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='注册Amazon MWS';

-- ----------------------------
-- Records of t_amazon_mws_config
-- ----------------------------

-- ----------------------------
-- Table structure for t_tenant
-- ----------------------------
DROP TABLE IF EXISTS `t_tenant`;
CREATE TABLE `t_tenant` (
  `tenant_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '租住名称',
  `organization` varchar(128) DEFAULT '' COMMENT '组织机构',
  `registered_address` varchar(200) DEFAULT '' COMMENT '注册地址',
  `mobile` varchar(255) DEFAULT NULL,
  `status` int(255) NOT NULL DEFAULT 1,
  `create_by` varchar(50) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(50) DEFAULT '',
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='平台租户';

-- ----------------------------
-- Records of t_tenant
-- ----------------------------
INSERT INTO `t_tenant` VALUES ('1', '深圳**有限公司', null, null, '13760330153', '0', 'test', '2021-02-24 15:28:28', 'test', '2021-02-24 15:28:24');

-- ----------------------------
-- Table structure for t_users
-- ----------------------------
DROP TABLE IF EXISTS `t_users`;
CREATE TABLE `t_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` int(11) NOT NULL DEFAULT 0 COMMENT '租户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(32) NOT NULL COMMENT '密码MD5',
  `status` int(255) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(50) DEFAULT '' COMMENT '修改人',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_username` (`username`) USING BTREE,
  KEY `username` (`username`),
  KEY `tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of t_users
-- ----------------------------
