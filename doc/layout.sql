/*
 Navicat Premium Dump SQL

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 80407 (8.4.7)
 Source Host           : localhost:3306
 Source Schema         : layout

 Target Server Type    : MySQL
 Target Server Version : 80407 (8.4.7)
 File Encoding         : 65001

 Date: 30/11/2025 20:48:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `user_id` bigint DEFAULT '0' COMMENT '负责人',
  `phone` varchar(64) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(64) DEFAULT NULL COMMENT '邮件',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:dept` (`tenant_id`,`deleted`,`status`,`parent_id`,`name`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`id`, `name`, `parent_id`, `sort`, `user_id`, `phone`, `email`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统', 0, 0, 1, NULL, NULL, 0, 0, 1, NULL, '2025-11-22 09:15:04', NULL, '2025-11-22 09:15:08');
INSERT INTO `sys_dept` (`id`, `name`, `parent_id`, `sort`, `user_id`, `phone`, `email`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '测试', 1, 0, 1, NULL, NULL, 0, 0, 1, 'admin', '2025-11-25 14:15:18', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `label` varchar(64) NOT NULL COMMENT '标签',
  `value` varchar(64) NOT NULL COMMENT '键值',
  `dict_type_id` bigint NOT NULL COMMENT '字典类型',
  `color_type` varchar(64) DEFAULT NULL COMMENT '颜色类型',
  `css_class` varchar(64) DEFAULT NULL COMMENT 'CSS样式',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:dict` (`dict_type_id`,`status`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典';

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 0, '正常', '0', 1, NULL, NULL, 0, NULL, NULL, '2025-11-21 11:02:04', NULL, '2025-11-21 11:02:04');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, 1, '停用', '1', 1, NULL, NULL, 0, NULL, NULL, '2025-11-21 11:02:21', NULL, '2025-11-21 11:02:21');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (3, 0, '否', '0', 2, 'default', NULL, 0, NULL, '', '2025-11-21 14:26:37', '', '2025-11-21 14:30:59');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (4, 1, '是', '1', 2, 'danger', NULL, 0, NULL, '', '2025-11-21 14:27:10', '', '2025-11-21 14:31:16');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (5, 0, '目录', '0', 6, 'primary', NULL, 0, NULL, '', '2025-11-21 14:29:13', '', '2025-11-21 16:05:36');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, 0, '菜单', '1', 6, 'success', NULL, 0, NULL, '', '2025-11-21 14:29:27', '', '2025-11-21 16:05:59');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, 0, '按钮', '2', 6, 'info', NULL, 0, NULL, '', '2025-11-21 14:29:35', '', '2025-11-21 16:06:08');
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, 0, '是', '1', 7, NULL, NULL, 0, NULL, '', '2025-11-21 14:32:18', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (9, 0, '否', '0', 7, NULL, NULL, 0, NULL, '', '2025-11-21 14:32:28', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (10, 0, '是', '1', 8, NULL, NULL, 0, NULL, '', '2025-11-21 14:33:31', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (11, 0, '否', '0', 8, NULL, NULL, 0, NULL, '', '2025-11-21 14:33:38', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (12, 0, '是', '1', 9, NULL, NULL, 0, NULL, '', '2025-11-21 14:34:25', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (13, 0, '否', '0', 9, NULL, NULL, 0, NULL, '', '2025-11-21 14:34:33', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (14, 1, '全部数据权限', '1', 10, NULL, NULL, 0, NULL, 'admin', '2025-11-25 21:30:44', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (15, 2, '自定数据权限', '2', 10, NULL, NULL, 0, NULL, 'admin', '2025-11-25 21:31:06', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (16, 3, '本部门数据权限', '3', 10, NULL, NULL, 0, NULL, 'admin', '2025-11-25 21:31:24', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (17, 4, '本部门及以下数据权限', '4', 10, NULL, NULL, 0, NULL, 'admin', '2025-11-25 21:31:40', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (18, 5, '仅自己', '5', 10, NULL, NULL, 0, NULL, 'admin', '2025-11-25 21:32:00', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (19, 0, 'POST', 'POST', 11, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:30:07', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, 0, 'PUT', 'PUT', 11, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:30:20', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (21, 0, 'DELETE', 'DELETE', 11, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:30:33', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (22, 0, 'GET', 'GET', 11, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:30:46', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (23, 0, '成功', '0', 12, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:31:36', NULL, NULL);
INSERT INTO `sys_dict` (`id`, `sort`, `label`, `value`, `dict_type_id`, `color_type`, `css_class`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (24, 0, '失败', '1', 12, NULL, NULL, 0, NULL, 'admin', '2025-11-29 09:31:45', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '字典名称',
  `type` varchar(64) NOT NULL COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:type:dicttype` (`type`),
  KEY `idx:list:dicttype` (`status`,`type`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '状态', 'status', 0, '系统状态', NULL, '2025-11-21 11:01:29', 'admin', '2025-11-25 12:56:50');
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '删除', 'deleted', 0, '删除', '', '2025-11-21 11:09:57', '', '2025-11-21 11:20:29');
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, '菜单类型', 'menu.type', 0, '', '', '2025-11-21 11:21:40', '', '2025-11-21 14:28:42');
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, '菜单隐藏', 'menu.hide', 0, NULL, '', '2025-11-21 14:31:56', NULL, NULL);
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, '菜单全屏', 'menu.full', 0, NULL, '', '2025-11-21 14:33:19', NULL, NULL);
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (9, '菜单缓存', 'menu.cache', 0, NULL, '', '2025-11-21 14:34:10', NULL, NULL);
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (10, '角色数据范围', 'role.scope', 0, NULL, 'admin', '2025-11-25 21:30:20', NULL, NULL);
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (11, '操作日志请求方法', 'operate.method', 0, NULL, 'admin', '2025-11-29 09:29:11', NULL, NULL);
INSERT INTO `sys_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (12, '操作日志结果', 'operate.result', 0, NULL, 'admin', '2025-11-29 09:31:19', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_logger_dev
-- ----------------------------
DROP TABLE IF EXISTS `sys_logger_dev`;
CREATE TABLE `sys_logger_dev` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `host` varchar(64) DEFAULT NULL COMMENT '服务名',
  `timestamp` datetime NOT NULL COMMENT '时间',
  `file` varchar(255) DEFAULT NULL,
  `func` varchar(255) DEFAULT NULL COMMENT '方法名',
  `message` varchar(512) DEFAULT NULL COMMENT '消息',
  `level` varchar(64) DEFAULT NULL COMMENT '等级',
  `data` json DEFAULT NULL COMMENT '数据',
  PRIMARY KEY (`id`),
  KEY `idx:list:dev` (`timestamp`,`host`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='开发日志';

-- ----------------------------
-- Records of sys_logger_dev
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_logger_login
-- ----------------------------
DROP TABLE IF EXISTS `sys_logger_login`;
CREATE TABLE `sys_logger_login` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) DEFAULT NULL COMMENT '姓名',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `ua` varchar(2048) DEFAULT NULL COMMENT 'UA',
  `login_time` datetime NOT NULL COMMENT '登录时间',
  `channel` varchar(64) DEFAULT NULL COMMENT '渠道',
  `ip` varchar(64) DEFAULT NULL COMMENT 'IP',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:login` (`tenant_id`,`deleted`,`channel`,`login_time`,`user_id`,`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志';

-- ----------------------------
-- Records of sys_logger_login
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_logger_operate
-- ----------------------------
DROP TABLE IF EXISTS `sys_logger_operate`;
CREATE TABLE `sys_logger_operate` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) DEFAULT NULL COMMENT '姓名',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `module` varchar(64) DEFAULT NULL COMMENT '模块名称',
  `method` varchar(64) DEFAULT NULL COMMENT '请求方法',
  `url` varchar(64) DEFAULT NULL COMMENT '请求地址',
  `ip` varchar(64) DEFAULT NULL COMMENT 'IP',
  `ua` varchar(2048) DEFAULT NULL COMMENT 'UA',
  `go_method` varchar(64) DEFAULT NULL COMMENT '方法名',
  `go_method_args` json DEFAULT NULL COMMENT '方法参数',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `duration` int DEFAULT '0' COMMENT '执行时长',
  `channel` varchar(64) DEFAULT NULL COMMENT '渠道',
  `result` tinyint DEFAULT '0' COMMENT '结果:0 成功/1 失败',
  `result_msg` varchar(255) DEFAULT NULL COMMENT '结果信息',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:operate` (`tenant_id`,`deleted`,`result`,`channel`,`method`,`start_time`,`user_id`,`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

-- ----------------------------
-- Records of sys_logger_operate
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `code` varchar(64) DEFAULT NULL COMMENT '编码',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '类型:0 目录/1 菜单/2 按钮',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级',
  `path` varchar(128) DEFAULT NULL COMMENT '地址',
  `icon` varchar(64) DEFAULT NULL COMMENT '图标',
  `component` varchar(128) DEFAULT NULL COMMENT '组件路径',
  `component_name` varchar(64) DEFAULT NULL COMMENT '组件名称',
  `hide` tinyint NOT NULL DEFAULT '0' COMMENT '隐藏:0 否/1 是',
  `link` varchar(128) DEFAULT NULL COMMENT '外部地址',
  `cache` tinyint NOT NULL DEFAULT '1' COMMENT '缓存:0否/1 是',
  `remark` varchar(64) DEFAULT NULL COMMENT '备注',
  `active` varchar(128) DEFAULT NULL COMMENT '激活地址',
  `full` tinyint NOT NULL DEFAULT '0' COMMENT '全屏:1 开/0 关',
  `redirect` varchar(128) DEFAULT NULL COMMENT '重定向',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `list:parent:menu` (`status`,`type`,`parent_id`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统管理', NULL, 0, 1, 0, '/sys', 'ep:setting', NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 15:56:50', 'admin', '2025-11-22 20:03:42');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '菜单管理', 'menu.SysMenuList', 1, 1, 1, '/sys/menu', 'ep:menu', '/sys/menu/index', 'SysMenu', 0, NULL, 1, NULL, NULL, 0, NULL, 0, '', '2025-11-21 15:59:34', 'admin', '2025-11-26 11:48:31');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (3, '菜单编辑', 'menu.SysMenuUpdate', 2, 0, 2, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:08:20', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (4, '菜单新增', 'menu.SysMenuCreate', 2, 0, 2, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:08:52', NULL, '2025-11-30 09:45:52');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (5, '菜单删除', 'menu.SysMenuDelete', 2, 0, 2, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:09:13', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, '字典管理', 'dict.SysDictTypeList', 1, 10, 1, '/sys/dict', 'ep:collection', '/sys/dict/index', 'SysDictType', 0, NULL, 1, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:14:37', 'admin', '2025-11-26 11:49:38');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, '字典类型新增', 'dict.SysDictTypeCreate', 2, 2, 6, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:17:26', 'admin', '2025-11-23 08:01:52');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, '字典类型编辑', 'dict.SysDictTypeUpdate', 2, 3, 6, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:18:23', 'admin', '2025-11-23 08:02:03');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (9, '字典类型删除', 'dict.SysDictTypeDelete', 2, 1, 6, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:18:45', 'admin', '2025-11-23 08:01:43');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (10, '字典列表', 'dict.SysDictList', 1, 0, 6, '/sys/dict/:dictTypeId', 'ep:data-board', '/sys/dict/item', 'SysDict', 1, NULL, 1, NULL, '/sys/dict', 0, NULL, 0, '', '2025-11-21 16:20:07', 'admin', '2025-11-30 11:06:38');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (11, '字典新增', 'dict.SysDictCreate', 2, 0, 10, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (12, '字典编辑', 'dict.SysDictUpdate', 2, 0, 10, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (13, '字典删除', 'dict.SysDictDelete', 2, 0, 10, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (16, '菜单查看', 'menu.SysMenu', 2, 0, 2, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (17, '字典查看', 'dict.SysDict', 2, 0, 10, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:38:14', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (18, '字典类型查看', 'dict.SysDictType', 2, 4, 6, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:38:38', 'admin', '2025-11-24 10:38:51');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (19, '租户管理', NULL, 0, 2, 1, '/tenant', 'ip:every-user', NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 17:06:39', 'admin', '2025-11-24 21:56:51');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, '租户列表', 'tenant.SysTenantList', 1, 0, 19, '/sys/tenant', 'ep:user', '/sys/tenant/index', 'SysTenant', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 17:10:29', 'admin', '2025-11-26 11:48:42');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (21, '租户服务', 'tenant.SysTenantPackageList', 1, 1, 19, '/sys/tenant/package', 'ip:activity-source', '/sys/tenant/package', 'SysTenantPackage', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 17:14:46', 'admin', '2025-11-26 11:48:49');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (22, '租户新增', 'tenant.SysTenantCreate', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 17:17:09', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (23, '租户编辑', 'tenant.SysTenantUpdate', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (24, '租户删除', 'tenant.SysTenantDelete', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (25, '租户查看', 'tenant.SysTenant', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (26, '租户清理', 'tenant.SysTenantDrop', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (27, '租户恢复', 'tenant.SysTenantRecover', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', 'admin', '2025-11-24 17:24:24');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (28, '租户服务新增', 'tenant.SysTenantPackageCreate', 2, 0, 21, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (29, '租户服务编辑', 'tenant.SysTenantPackageUpdate', 2, 0, 21, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (30, '租户服务删除', 'tenant.SysTenantPackageDelete', 2, 0, 21, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (31, '租户服务查看', 'tenant.SysTenantPackage', 2, 0, 21, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (32, '职位管理', 'post.SysPostList', 1, 3, 1, '/sys/post', 'ep:postcard', '/sys/post/index', NULL, 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 22:53:14', 'admin', '2025-11-26 11:48:59');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (33, '职位新增', 'post.SysPostCreate', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', 'admin', '2025-11-24 22:57:56');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (34, '职位编辑', 'post.SysPostUpdate', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (35, '职位删除', 'post.SysPostDelete', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (36, '职位查看', 'post.SysPost', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (37, '职位清理', 'post.SysPostDrop', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (38, '职位恢复', 'post.SysPostRecover', 2, 0, 32, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (39, '部门管理', 'dept.SysDeptList', 1, 4, 1, '/sys/dept', 'ip:tree-list', '/sys/dept/index', 'SysDept', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-25 13:57:18', 'admin', '2025-11-26 11:49:09');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (40, '部门新增', 'dept.SysDeptCreate', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (41, '部门编辑', 'dept.SysDeptUpdate', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (42, '部门删除', 'dept.SysDeptDelete', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (43, '部门查看', 'dept.SysDept', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (44, '部门清理', 'dept.SysDeptDrop', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (45, '部门恢复', 'dept.SysDeptRecover', 2, 0, 39, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (46, '用户管理', 'user.SysUserList', 1, 5, 1, '/sys/user', 'ip:user-business', '/sys/user/index', 'SysUser', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-25 14:34:37', 'admin', '2025-11-26 11:49:25');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (47, '用户新增', 'user.SysUserCreate', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (48, '用户编辑', 'user.SysUserUpdate', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (49, '用户删除', 'user.SysUserDelete', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (50, '用户查看', 'user.SysUser', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (51, '用户清理', 'user.SysUserDrop', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (52, '用户恢复', 'user.SysUserRecover', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (53, '角色管理', 'role.SysRoleList', 1, 6, 1, '/sys/role', 'ip:user-to-user-transmission', '/sys/role/index', 'SysRole', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-25 21:15:10', 'admin', '2025-11-26 13:53:13');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (54, '角色新增', 'role.SysRoleCreate', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:20:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (55, '角色编辑', 'role.SysRoleUpdate', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:21:04', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (56, '角色删除', 'role.SysRoleDelete', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, '', '2025-11-21 16:22:01', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (57, '角色查看', 'role.SysRole', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 20:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (58, '角色清理', 'role.SysRoleDrop', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (59, '角色恢复', 'role.SysRoleRecover', 2, 0, 53, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-24 10:37:41', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (60, '日志管理', NULL, 0, 7, 1, '/sys/logger', 'ip:log', NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 13:53:58', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (61, '开发日志', 'logger.SysLoggerDevList', 1, 0, 60, '/sys/logger/dev', 'ip:coordinate-system', '/sys/logger/dev', 'SysLoggerDev', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 13:55:32', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (62, '开发日志查看', 'logger.SysLoggerDev', 2, 0, 61, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 13:56:07', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (63, '开发日志删除', 'logger.SysLoggerDevDelete', 2, 0, 61, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 13:56:37', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (64, '登录日志', 'logger.SysLoggerLoginList', 1, 0, 60, '/sys/logger/login', 'ip:login', '/sys/logger/login', 'SysLoggerLogin', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 13:59:21', 'admin', '2025-11-30 11:05:27');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (65, '操作日志', 'logger.SysLoggerOperateList', 1, 0, 60, '/sys/logger/operate', 'ep:document-copy', '/sys/logger/operate', 'SysLoggerOperate', 0, NULL, 1, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-26 14:00:06', 'admin', '2025-11-30 11:05:19');
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (66, '操作日志查看', 'logger.SysLoggerOperate', 2, 0, 65, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 09:18:24', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (67, '操作日志删除', 'logger.SysLoggerOperateDelete', 2, 0, 65, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 09:19:00', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (68, '操作日志清理', 'logger.SysLoggerOperateDrop', 2, 0, 65, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 09:19:25', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (69, '操作日志恢复', 'logger.SysLoggerOperateRecover', 2, 0, 65, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 09:19:48', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (70, '用户登录', 'user.SysUserLogin', 2, 0, 46, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 13:35:48', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (71, '租户登录', 'tenant.SysTenantUserLogin', 2, 0, 20, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-29 13:36:19', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (72, '登录日志查看', 'logger.SysLoggerLogin', 2, 0, 64, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-30 09:47:00', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (73, '登录日志删除', 'logger.SysLoggerLoginDelete', 2, 0, 64, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-30 09:47:34', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (74, '登录日志清理', 'logger.SysLoggerLoginDrop', 2, 0, 64, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-30 09:47:59', NULL, NULL);
INSERT INTO `sys_menu` (`id`, `name`, `code`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `hide`, `link`, `cache`, `remark`, `active`, `full`, `redirect`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (75, '登录日志恢复', 'logger.SysLoggerLoginRecover', 2, 0, 64, NULL, NULL, NULL, NULL, 0, NULL, 0, NULL, NULL, 0, NULL, 0, 'admin', '2025-11-30 09:48:20', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:post` (`tenant_id`,`deleted`,`status`,`name`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='职位';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`id`, `name`, `sort`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '销售', 0, 0, 0, 1, NULL, '2025-11-04 20:14:35', 'admin', '2025-11-25 23:39:57');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `code` varchar(64) NOT NULL COMMENT '编码',
  `scope` tinyint DEFAULT NULL COMMENT '数据范围:1:全部数据权限/2:自定数据权限/3:本部门数据权限/4:本部门及以下数据权限',
  `scope_dept` json DEFAULT NULL COMMENT '数据范围(指定部门数组)',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:code:role` (`tenant_id`,`code`),
  KEY `idx:list:role` (`tenant_id`,`deleted`,`status`,`code`,`name`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `name`, `code`, `scope`, `scope_dept`, `sort`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, 'asd', 'asd', 2, '[1, 2]', 0, 0, 0, 1, 'admin', '2025-11-25 23:05:52', 'admin', '2025-11-26 10:10:42');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `role_id` bigint NOT NULL COMMENT '角色',
  `menu_id` bigint NOT NULL COMMENT '菜单',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:item:role` (`tenant_id`,`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=596 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (539, 20, 1, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (540, 20, 2, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (544, 20, 3, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (543, 20, 4, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (542, 20, 5, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (586, 20, 6, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (593, 20, 7, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (594, 20, 8, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (592, 20, 9, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (587, 20, 10, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (591, 20, 11, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (590, 20, 12, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (589, 20, 13, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (541, 20, 16, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (588, 20, 17, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (595, 20, 18, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (545, 20, 19, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (546, 20, 20, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (553, 20, 21, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (552, 20, 22, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (551, 20, 23, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (550, 20, 24, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (549, 20, 25, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (548, 20, 26, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (547, 20, 27, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (557, 20, 28, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (556, 20, 29, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (555, 20, 30, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (554, 20, 31, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (558, 20, 32, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (564, 20, 33, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (563, 20, 34, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (562, 20, 35, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (561, 20, 36, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (560, 20, 37, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (559, 20, 38, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (565, 20, 39, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (571, 20, 40, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (570, 20, 41, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (569, 20, 42, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (568, 20, 43, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (567, 20, 44, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (566, 20, 45, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (572, 20, 46, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (578, 20, 47, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (577, 20, 48, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (576, 20, 49, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (575, 20, 50, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (574, 20, 51, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (573, 20, 52, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (579, 20, 53, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (585, 20, 54, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (584, 20, 55, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (583, 20, 56, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (582, 20, 57, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (581, 20, 58, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `tenant_id`) VALUES (580, 20, 59, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `user_id` bigint DEFAULT '0' COMMENT '用户',
  `contact_name` varchar(64) DEFAULT NULL COMMENT '联系人',
  `contact_mobile` varchar(64) DEFAULT NULL COMMENT '联系电话',
  `expire_date` date NOT NULL COMMENT '过期时间',
  `account_total` bigint NOT NULL DEFAULT '0' COMMENT '账号数量',
  `package_id` bigint NOT NULL DEFAULT '0' COMMENT '套餐编号',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:package:tenant` (`package_id`),
  KEY `idx:list:tenant` (`deleted`,`status`,`name`,`expire_date`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户';

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sys_tenant` (`id`, `name`, `user_id`, `contact_name`, `contact_mobile`, `expire_date`, `account_total`, `package_id`, `status`, `deleted`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统', 1, 'admin', '13888888888', '2026-09-30', 10000, 1, 0, 0, NULL, '2025-11-22 09:13:22', 'admin', '2025-11-25 12:56:21');
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant_package
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant_package`;
CREATE TABLE `sys_tenant_package` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `scope_menu` json DEFAULT NULL COMMENT '数据范围',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list:package` (`status`,`name`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户套餐';

-- ----------------------------
-- Records of sys_tenant_package
-- ----------------------------
BEGIN;
INSERT INTO `sys_tenant_package` (`id`, `name`, `scope_menu`, `sort`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统', '[1, 2, 16, 5, 4, 3, 19, 20, 71, 27, 26, 25, 24, 23, 22, 21, 31, 30, 29, 28, 32, 38, 37, 36, 35, 34, 33, 39, 45, 44, 43, 42, 41, 40, 46, 70, 52, 51, 50, 49, 48, 47, 53, 59, 58, 57, 56, 55, 54, 60, 65, 69, 68, 67, 66, 64, 75, 74, 73, 72, 61, 63, 62, 6, 10, 17, 13, 12, 11, 9, 7, 8, 18]', 0, 0, '系统', NULL, '2025-11-22 09:14:33', 'admin', '2025-11-30 09:48:53');
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) DEFAULT NULL COMMENT '姓名',
  `mobile` varchar(64) DEFAULT NULL COMMENT '手机号码',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除:0否/1是',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:login:user` (`username`),
  KEY `idx:list:user` (`tenant_id`,`deleted`,`status`,`mobile`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `name`, `mobile`, `username`, `password`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 'adminwww', '13888888888', 'admin', '21232f297a57a5a743894a0e4a801fc3', 0, 0, 1, NULL, '2025-11-22 09:12:27', 'admin', '2025-11-26 10:45:22');
INSERT INTO `sys_user` (`id`, `name`, `mobile`, `username`, `password`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (3, 'asdas', 'asd', 'asd', 'asd', 0, 0, 1, 'admin', '2025-11-26 11:24:48', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_dept`;
CREATE TABLE `sys_user_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户',
  `dept_id` bigint NOT NULL COMMENT '部门',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:item:user` (`tenant_id`,`user_id`,`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户部门';

-- ----------------------------
-- Records of sys_user_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_dept` (`id`, `user_id`, `dept_id`, `tenant_id`) VALUES (3, 1, 1, 1);
INSERT INTO `sys_user_dept` (`id`, `user_id`, `dept_id`, `tenant_id`) VALUES (4, 1, 2, 1);
INSERT INTO `sys_user_dept` (`id`, `user_id`, `dept_id`, `tenant_id`) VALUES (7, 3, 1, 1);
INSERT INTO `sys_user_dept` (`id`, `user_id`, `dept_id`, `tenant_id`) VALUES (8, 3, 2, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户',
  `post_id` bigint NOT NULL COMMENT '部门',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:item:post` (`tenant_id`,`user_id`,`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户职位';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` (`id`, `user_id`, `post_id`, `tenant_id`) VALUES (2, 1, 2, 1);
INSERT INTO `sys_user_post` (`id`, `user_id`, `post_id`, `tenant_id`) VALUES (4, 3, 2, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户',
  `role_id` bigint NOT NULL COMMENT '角色',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  PRIMARY KEY (`id`),
  KEY `item:item:role` (`tenant_id`,`user_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`, `tenant_id`) VALUES (2, 1, 20, 1);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`, `tenant_id`) VALUES (4, 3, 20, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_tenant`;
CREATE TABLE `sys_user_tenant` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:bind:user` (`user_id`,`tenant_id`),
  KEY `idx:tenant:user` (`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户用户';

-- ----------------------------
-- Records of sys_user_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_tenant` (`id`, `user_id`, `tenant_id`) VALUES (1, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
