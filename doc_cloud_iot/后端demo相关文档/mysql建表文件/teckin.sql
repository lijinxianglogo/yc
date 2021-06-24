/*
 Navicat Premium Data Transfer

 Source Server         : 海外开发环境
 Source Server Type    : MySQL
 Source Server Version : 50712
 Source Host           : ap-iot-test-1.cluster-cvrzlz5zm10t.ap-northeast-1.rds.amazonaws.com:3306
 Source Schema         : teckin

 Target Server Type    : MySQL
 Target Server Version : 50712
 File Encoding         : 65001

 Date: 14/05/2021 16:59:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_device
-- ----------------------------
DROP TABLE IF EXISTS `t_device`;
CREATE TABLE `t_device`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `code` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid Éè±¸Î¨Ò»Ê¶±ðÂë',
  `name` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `model` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `version` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_time` int(11) NOT NULL COMMENT 'Ãë¼¶Ê±¼ä´Á',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `t_device_code`(`code`) USING BTREE,
  INDEX `FK_user_dev_relate`(`user_id`) USING BTREE,
  CONSTRAINT `FK_user_dev_relate` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Ö÷¼üID',
  `code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'ÓÃÓÚÊ¶±ðÓÃ»§µÄÐÅÏ¢',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `mobile` int(11) NULL DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `idcard` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sex` smallint(6) NULL DEFAULT NULL COMMENT '1:ÄÐ£¬2£ºÅ®',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `add_time` bigint(20) NOT NULL COMMENT 'Ãë¼¶Ê±¼ä´Á',
  `last_time` bigint(20) NULL DEFAULT NULL COMMENT 'Ãë¼¶Ê±¼ä´Á',
  `is_del` smallint(6) NOT NULL COMMENT '0£º×¢Ïú£¬1£ºÎ´×¢Ïú',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_userquerule
-- ----------------------------
DROP TABLE IF EXISTS `t_userquerule`;
CREATE TABLE `t_userquerule`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_u_id` int(11) NOT NULL COMMENT 'Ö÷¼üID',
  `topic` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_user_rule_relate`(`t_u_id`) USING BTREE,
  CONSTRAINT `FK_user_rule_relate` FOREIGN KEY (`t_u_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
