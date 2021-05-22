grant all privileges on dbname.tablename to 'chat'@'%';
grant all privileges on chat.* to 'chat'@'%';
flush privileges;

-- 创建数据库
CREATE DATABASE IF NOT EXISTS `chat`;
USE `chat`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apply
-- ----------------------------
DROP TABLE IF EXISTS `apply`;
CREATE TABLE `apply` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `friend_id` int(11) unsigned NOT NULL COMMENT '好友id',
  `nickname` varchar(60) NOT NULL COMMENT '备注昵称',
  `look_me` tinyint(4) NOT NULL DEFAULT '1' COMMENT '看我',
  `look_him` tinyint(4) NOT NULL DEFAULT '1' COMMENT '看他',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_apply_friend_id` (`friend_id`),
  KEY `idx_apply_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for collect
-- ----------------------------
DROP TABLE IF EXISTS `collect`;
CREATE TABLE `collect` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `content` varchar(5000) NOT NULL COMMENT '内容',
  `type` tinyint(4) NOT NULL COMMENT '类型',
  `options` varchar(255) NOT NULL DEFAULT '' COMMENT '选项',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_collect_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `friend_id` int(11) unsigned NOT NULL COMMENT '好友id',
  `nickname` varchar(60) NOT NULL COMMENT '备注昵称',
  `look_me` tinyint(4) NOT NULL DEFAULT '1' COMMENT '看我',
  `look_him` tinyint(4) NOT NULL DEFAULT '1' COMMENT '看他',
  `is_star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否星标用户',
  `is_black` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否拉黑',
  `tags` varchar(1000) NOT NULL DEFAULT '' COMMENT '标签',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uid_fid` (`user_id`,`friend_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `name` varchar(255) NOT NULL COMMENT '群组名',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `invite_confirm` tinyint(4) NOT NULL DEFAULT '0' COMMENT '邀请确认',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for group_user
-- ----------------------------
DROP TABLE IF EXISTS `group_user`;
CREATE TABLE `group_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `group_id` int(11) unsigned NOT NULL COMMENT '群组ID',
  `nickname` varchar(60) NOT NULL COMMENT '备注昵称',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_user_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `to_id` int(11) unsigned NOT NULL COMMENT '发送者',
  `chat_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '目标类型，1=用户，2=群组',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '消息类型',
  `content` varchar(5000) NOT NULL COMMENT '内容',
  `options` varchar(255) NOT NULL DEFAULT '' COMMENT '选项',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_message_user_id` (`user_id`),
  KEY `idx_message_to_id` (`to_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for moment
-- ----------------------------
DROP TABLE IF EXISTS `moment`;
CREATE TABLE `moment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `content` varchar(1000) NOT NULL COMMENT '内容',
  `image` varchar(1000) NOT NULL DEFAULT '' COMMENT '图片',
  `video` varchar(255) NOT NULL DEFAULT '' COMMENT '视频地址',
  `location` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `remind` varchar(255) NOT NULL DEFAULT '' COMMENT '提醒谁看',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '动态类型',
  `see_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '可见类型',
  `see` varchar(255) NOT NULL DEFAULT '' COMMENT '用户id列表',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_moment_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for moment_comment
-- ----------------------------
DROP TABLE IF EXISTS `moment_comment`;
CREATE TABLE `moment_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `reply_id` int(11) unsigned NOT NULL COMMENT '回复用户id',
  `moment_id` int(11) unsigned NOT NULL COMMENT '动态id',
  `content` varchar(1000) NOT NULL COMMENT '评论内容',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_moment_comment_user_id` (`user_id`),
  KEY `idx_moment_comment_reply_id` (`reply_id`),
  KEY `idx_moment_comment_moment_id` (`moment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for moment_like
-- ----------------------------
DROP TABLE IF EXISTS `moment_like`;
CREATE TABLE `moment_like` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `moment_id` int(11) unsigned NOT NULL COMMENT '动态id',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uid_mid` (`user_id`,`moment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for moment_timeline
-- ----------------------------
DROP TABLE IF EXISTS `moment_timeline`;
CREATE TABLE `moment_timeline` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `moment_id` int(11) unsigned NOT NULL COMMENT '动态id',
  `is_own` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否自己的',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_moment_timeline_user_id` (`user_id`),
  KEY `idx_moment_timeline_moment_id` (`moment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for report
-- ----------------------------
DROP TABLE IF EXISTS `report`;
CREATE TABLE `report` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `target_id` int(11) unsigned NOT NULL COMMENT '目标id',
  `target_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '目标类型，1=用户，2=群组',
  `content` varchar(5000) NOT NULL COMMENT '内容',
  `category` varchar(255) NOT NULL DEFAULT '' COMMENT '分类',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_report_user_id` (`user_id`),
  KEY `idx_report_target_id` (`target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `phone` bigint(20) NOT NULL COMMENT '手机号',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(4) NOT NULL DEFAULT '1' COMMENT '性别',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `sign` varchar(255) NOT NULL DEFAULT '' COMMENT '签名',
  `area` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_username` (`username`),
  UNIQUE KEY `idx_user_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_tag
-- ----------------------------
DROP TABLE IF EXISTS `user_tag`;
CREATE TABLE `user_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `name` varchar(60) NOT NULL COMMENT '标签名',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`user_id`,`name`),
  KEY `idx_user_tag_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
