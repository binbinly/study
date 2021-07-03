/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.8.76
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.8.76:3306
 Source Schema         : chat_admin

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 25/06/2021 16:42:40
*/

-- 创建数据库
CREATE DATABASE IF NOT EXISTS `chat_admin`;
USE `chat_admin`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `permission` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, 1, 'Dashboard', 'fa-bar-chart', '/', NULL, NULL, NULL);
INSERT INTO `admin_menu` VALUES (2, 0, 3, 'admin', 'fa-tasks', '', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (3, 2, 4, 'Users', 'fa-users', 'auth/users', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (4, 2, 5, 'Roles', 'fa-user', 'auth/roles', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (5, 2, 6, 'Permission', 'fa-ban', 'auth/permissions', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (6, 2, 7, 'Menu', 'fa-bars', 'auth/menu', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (7, 2, 8, 'Operation log', 'fa-history', 'auth/logs', NULL, NULL, '2021-06-24 16:03:55');
INSERT INTO `admin_menu` VALUES (8, 0, 2, '表情包', 'fa-heart-o', '/emoticon', '', '2021-06-24 16:03:48', '2021-06-24 16:03:55');
COMMIT;

-- ----------------------------
-- Table structure for admin_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_operation_log`;
CREATE TABLE `admin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_operation_log
-- ----------------------------
BEGIN;
INSERT INTO `admin_operation_log` VALUES (1, 1, 'auth/setting', 'PUT', '127.0.0.1', '{\"name\":\"Administrator\",\"password\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"password_confirmation\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"_token\":\"TPiG7q7ma3my2KlCCJ6FQMVcTpza7A30odi9qrGe\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/admin.chat.lo\\/\"}', '2021-06-22 06:07:22', '2021-06-22 06:07:22');
INSERT INTO `admin_operation_log` VALUES (2, 1, 'auth/setting', 'PUT', '127.0.0.1', '{\"name\":\"Administrator\",\"password\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"password_confirmation\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"_token\":\"TPiG7q7ma3my2KlCCJ6FQMVcTpza7A30odi9qrGe\",\"_method\":\"PUT\"}', '2021-06-22 06:07:53', '2021-06-22 06:07:53');
INSERT INTO `admin_operation_log` VALUES (3, 1, 'auth/setting', 'PUT', '127.0.0.1', '{\"name\":\"Administrator\",\"password\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"password_confirmation\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"_token\":\"TPiG7q7ma3my2KlCCJ6FQMVcTpza7A30odi9qrGe\",\"_method\":\"PUT\"}', '2021-06-22 06:12:35', '2021-06-22 06:12:35');
INSERT INTO `admin_operation_log` VALUES (4, 1, 'auth/setting', 'PUT', '127.0.0.1', '{\"name\":\"Administrator\",\"password\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"password_confirmation\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"_token\":\"TPiG7q7ma3my2KlCCJ6FQMVcTpza7A30odi9qrGe\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/admin.chat.lo\\/auth\\/login\"}', '2021-06-22 06:23:13', '2021-06-22 06:23:13');
INSERT INTO `admin_operation_log` VALUES (5, 1, 'auth/setting', 'PUT', '127.0.0.1', '{\"name\":\"Administrator\",\"password\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"password_confirmation\":\"$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u\",\"_token\":\"TPiG7q7ma3my2KlCCJ6FQMVcTpza7A30odi9qrGe\",\"_method\":\"PUT\"}', '2021-06-22 06:25:13', '2021-06-22 06:25:13');
INSERT INTO `admin_operation_log` VALUES (6, 1, 'auth/menu', 'POST', '127.0.0.1', '{\"parent_id\":\"0\",\"title\":\"\\u8868\\u60c5\\u5305\",\"icon\":\"fa-heart-o\",\"uri\":\"\\/emoticon\",\"roles\":[\"\"],\"permission\":\"\",\"_token\":\"MLYIzffuPuMgt6GsTuVncs5OdkcsPFGqJtMD9tDy\"}', '2021-06-24 16:03:48', '2021-06-24 16:03:48');
INSERT INTO `admin_operation_log` VALUES (7, 1, 'auth/menu', 'POST', '127.0.0.1', '{\"_token\":\"MLYIzffuPuMgt6GsTuVncs5OdkcsPFGqJtMD9tDy\",\"_order\":\"[{\\\"id\\\":1},{\\\"id\\\":8},{\\\"id\\\":2,\\\"children\\\":[{\\\"id\\\":3},{\\\"id\\\":4},{\\\"id\\\":5},{\\\"id\\\":6},{\\\"id\\\":7}]}]\"}', '2021-06-24 16:03:55', '2021-06-24 16:03:55');
INSERT INTO `admin_operation_log` VALUES (8, 1, 'emoticon/1', 'PUT', '127.0.0.1', '{\"category\":\"\\u8d21\\u732e\\ud83c\\udde8\\ud83c\\uddf3BQB\",\"name\":\"0000\",\"_token\":\"MLYIzffuPuMgt6GsTuVncs5OdkcsPFGqJtMD9tDy\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/admin.chat.lo\\/emoticon?view=card&category=%E8%B4%A1%E7%8C%AE%F0%9F%87%A8%F0%9F%87%B3BQB\"}', '2021-06-24 16:24:14', '2021-06-24 16:24:14');
INSERT INTO `admin_operation_log` VALUES (9, 1, 'emoticon', 'POST', '127.0.0.1', '{\"category\":\"\\u8d21\\u732e\\ud83c\\udde8\\ud83c\\uddf3BQB\",\"name\":\"aaa\",\"_token\":\"MLYIzffuPuMgt6GsTuVncs5OdkcsPFGqJtMD9tDy\",\"_previous_\":\"http:\\/\\/admin.chat.lo\\/emoticon?view=card&category=%E8%B4%A1%E7%8C%AE%F0%9F%87%A8%F0%9F%87%B3BQB\"}', '2021-06-24 16:25:06', '2021-06-24 16:25:06');
COMMIT;

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`),
  UNIQUE KEY `admin_permissions_slug_unique` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_permissions` VALUES (1, 'All permission', '*', '', '*', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (2, 'Dashboard', 'dashboard', 'GET', '/', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (3, 'Login', 'auth.login', '', '/auth/login\r\n/auth/logout', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (4, 'User setting', 'auth.setting', 'GET,PUT', '/auth/setting', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (5, 'Auth management', 'auth.management', '', '/auth/roles\r\n/auth/permissions\r\n/auth/menu\r\n/auth/logs', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menu`;
CREATE TABLE `admin_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_menu` VALUES (1, 1, 2, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permissions`;
CREATE TABLE `admin_role_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_role_permissions_role_id_permission_id_index` (`role_id`,`permission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_permissions` VALUES (1, 1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_users`;
CREATE TABLE `admin_role_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_role_users_role_id_user_id_index` (`role_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_users` VALUES (1, 1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`),
  UNIQUE KEY `admin_roles_slug_unique` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles` VALUES (1, 'Administrator', 'administrator', '2021-06-22 05:35:51', '2021-06-22 05:35:51');
COMMIT;

-- ----------------------------
-- Table structure for admin_user_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_permissions`;
CREATE TABLE `admin_user_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_user_permissions_user_id_permission_id_index` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_user_permissions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(190) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users` VALUES (1, 'admin', '$2y$10$IX3BJ9t6H4te.0ye7PqtbOjdiGtdvntu1DLmhAeeBOV2.Wf5TiF3u', 'Administrator', 'images/10.jpeg', 'p5dLa1rDIum3LaBflWSy1gPevEjkcfBExqxB5qyV6oG6w7UqyRlarrvA2xpx', '2021-06-22 05:35:51', '2021-06-22 06:25:13');
COMMIT;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of migrations
-- ----------------------------
BEGIN;
INSERT INTO `migrations` VALUES (1, '2014_10_12_000000_create_users_table', 1);
INSERT INTO `migrations` VALUES (2, '2014_10_12_100000_create_password_resets_table', 1);
INSERT INTO `migrations` VALUES (3, '2016_01_04_173148_create_admin_tables', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
