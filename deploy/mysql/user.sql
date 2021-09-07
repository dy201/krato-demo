/*
 Navicat Premium Data Transfer

 Source Server         : 本地1
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost:3306
 Source Schema         : doyle_blog

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 07/09/2021 17:04:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` int(3) NULL DEFAULT 0,
  `email` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` int(3) NULL DEFAULT NULL,
  `role_id` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name`) USING BTREE,
  INDEX `idx_nickname`(`nickname`) USING BTREE,
  INDEX `idx_phone`(`phone`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_username`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2021-09-06 15:48:01.087', '2021-09-06 15:48:01.087', '2021-09-07 13:38:44.588', 'admin', '123456', '管理员', 1, 'admin@123.com', '18801020304', 22, 0);
INSERT INTO `user` VALUES (2, '2021-09-06 21:24:26.762', '2021-09-06 21:24:26.762', NULL, 'doyle', '123456', '小良同学', 1, 'doyle@123.com', '18801020304', 22, 0);
INSERT INTO `user` VALUES (3, '2021-09-07 13:48:34.987', '2021-09-07 13:48:34.987', NULL, 'dolore ad ex ut deserunt', 'nostrud deserunt', 'sunt sint', 93568605, 'dolore anim labore', 'deserunt adipisicing', 111, 0);
INSERT INTO `user` VALUES (4, '2021-09-07 13:51:15.388', '2021-09-07 13:51:15.388', NULL, 'quis', 'voluptate cupidatat in pariatur', 'minim deserunt dolore in', 1, 'test@132.com', 'Lorem mollit ex', 111, 0);
INSERT INTO `user` VALUES (5, '2021-09-07 16:49:34.493', '2021-09-07 16:49:34.493', NULL, 'dalao', '$pbkdf2-sha512$pLRFWdyZeJDG5rDt$a5396b61bab15cade0f14f44fc3163a3f085599109e4da2de43fedc99eb9c57f', 'dalao', 0, 'dalao@123.com', '13311112222', 23, 0);

SET FOREIGN_KEY_CHECKS = 1;
