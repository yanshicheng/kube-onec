USE `kube_onec`;

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                            `user_name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '用户姓名',
                            `account` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户账号，唯一标识',
                            `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户密码，需加密存储',
                            `icon` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像URL',
                            `mobile` CHAR(11) NOT NULL DEFAULT '' COMMENT '用户手机号',
                            `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户邮箱地址',
                            `work_number` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户工号',
                            `hire_date` DATE NOT NULL DEFAULT '1970-01-01' COMMENT '入职日期',
                            `is_change_password` BOOLEAN DEFAULT 0 COMMENT '是否需要重置密码，0 否，1 是',
                            `is_disabled` BOOLEAN DEFAULT 0 COMMENT '是否禁用，0 否，1 是',
                            `is_leave` BOOLEAN DEFAULT 0 COMMENT '是否离职，0 否，1 是',
                            `position_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '职位ID，关联职位表',
                            `organization_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '组织ID，关联组织表',
                            `last_login_time` DATETIME DEFAULT '1970-01-01 00:00:00' COMMENT '上次登录时间',
                            `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                            `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（NULL表示未删除）',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `uk_account` (`account`),
                            UNIQUE KEY `uk_mobile` (`mobile`),
                            UNIQUE KEY `uk_email` (`email`),
                            UNIQUE KEY `uk_work_number` (`work_number`),
                            INDEX `idx_organization_id` (`organization_id`),
                            INDEX `idx_position_id` (`position_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号信息表';

DROP TABLE IF EXISTS `sys_organization`;
CREATE TABLE `sys_organization` (
                                    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '团队名称',
                                    `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父级组织的 ID，根级为 NULL',
                                    `level` INT DEFAULT 0 COMMENT '组织层级，从 0 开始',
                                    `description` TEXT DEFAULT NULL COMMENT '组织描述',
                                    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                    `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                                    PRIMARY KEY (`id`),
                                    INDEX `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='组织表';

DROP TABLE IF EXISTS `sys_position`;
CREATE TABLE `sys_position` (
                                `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '职位名称',
                                `organization_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属组织 ID，关联 sys_organization 表',
                                `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                                PRIMARY KEY (`id`),
                                INDEX `idx_organization_id` (`organization_id`),
                                UNIQUE KEY `uk_name_organization_id` (`name`, `organization_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职位表';

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                            `role_name` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '角色名称',
                            `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                            `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                            `description` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
                            `create_by` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '创建人',
                            `update_by` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '更新人',
                            PRIMARY KEY (`id`),
                            UNIQUE INDEX `uk_role_name` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID，关联 sys_user 表',
                                 `role_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID，关联 sys_role 表',
                                 `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                 `delete_time` TIMESTAMP NULL DEFAULT NULL COMMENT '记录删除时间（NULL表示未删除）',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_role_id` (`role_id`),
                                 UNIQUE KEY `uk_user_role` (`user_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户与角色关联表';

DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                            `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父级菜单ID',
                            `name` VARCHAR(255) DEFAULT NULL COMMENT '路由名称',
                            `path` VARCHAR(255) DEFAULT NULL COMMENT '路由路径',
                            `component` VARCHAR(255) DEFAULT NULL COMMENT '组件路径',
                            `component_name` VARCHAR(255) DEFAULT NULL COMMENT '组件名称',
                            `redirect` VARCHAR(255) DEFAULT NULL COMMENT '重定向路径',
    -- Meta相关配置字段，对应RouteMeta
                            `title` VARCHAR(255) DEFAULT '' COMMENT '标题名称（一般用于菜单和标签页显示，多配合国际化）',
                            `icon` VARCHAR(255) DEFAULT '' COMMENT '图标（菜单/tab）',
                            `active_icon` VARCHAR(255) DEFAULT '' COMMENT '激活图标（菜单）',
                            `active_path` VARCHAR(255) DEFAULT '' COMMENT '当前激活的菜单路径',

                            `affix_tab` BOOLEAN DEFAULT 0 COMMENT '是否固定标签页 0否1是',
                            `affix_tab_order` INT DEFAULT 0 COMMENT '固定标签页的排序',

                            `authority` JSON DEFAULT NULL COMMENT '访问权限标识数组，JSON存储[string, string...]',

                            `badge` VARCHAR(255) DEFAULT '' COMMENT '徽标文字',
                            `badge_type` ENUM('dot','normal') DEFAULT 'normal' COMMENT '徽标类型 dot|normal',
                            `badge_variants` VARCHAR(50) DEFAULT 'success' COMMENT '徽标颜色，如success,primary,warning等',

                            `hide_children_in_menu` BOOLEAN DEFAULT 0 COMMENT '是否在菜单中隐藏子路由 0否1是',
                            `hide_in_breadcrumb` BOOLEAN DEFAULT 0 COMMENT '是否在面包屑中隐藏 0否1是',
                            `hide_in_menu` BOOLEAN DEFAULT 0 COMMENT '是否在菜单中隐藏 0否1是',
                            `hide_in_tab` BOOLEAN DEFAULT 0 COMMENT '是否在标签页中隐藏 0否1是',

                            `iframe_src` VARCHAR(255) DEFAULT '' COMMENT 'iframe内嵌页面地址',
                            `ignore_access` BOOLEAN DEFAULT 0 COMMENT '是否忽略权限 0否1是',
                            `keep_alive` BOOLEAN DEFAULT 0 COMMENT '是否开启缓存 0否1是',

                            `link` VARCHAR(255) DEFAULT '' COMMENT '外链地址',

                            `loaded` BOOLEAN DEFAULT 0 COMMENT '路由是否已加载过(可选)',

                            `max_num_of_open_tab` INT DEFAULT -1 COMMENT '标签页最大打开数量，-1表示无限制',

                            `menu_visible_with_forbidden` BOOLEAN DEFAULT 0 COMMENT '菜单可见但访问403 0否1是',

                            `open_in_new_window` BOOLEAN DEFAULT 0 COMMENT '外链在新窗口打开 0否1是',

                            `order_num` INT DEFAULT 0 COMMENT '菜单排序，仅一级菜单有效',

                            `query` JSON DEFAULT NULL COMMENT '菜单参数JSON存储，如{"id":"123"}',
    -- 原有字段可根据需要保留或移除
                            `is_route` BOOLEAN DEFAULT 1 COMMENT '是否路由菜单: 0否1是',
                            `menu_type` INT DEFAULT 0 COMMENT '菜单类型(0:一级菜单;1:子菜单;2:按钮权限)',
                            `perms` VARCHAR(255) DEFAULT NULL COMMENT '菜单权限编码（旧项目字段，可根据需要保留）',
                            `perms_type` VARCHAR(10) DEFAULT '0' COMMENT '权限策略: 1显示2禁用（旧项目字段）',
                            `always_show` BOOLEAN DEFAULT NULL COMMENT '是否聚合子路由:1是0否（旧项目字段）',

                            `rule_flag` BOOLEAN DEFAULT 0 COMMENT '数据权限标记 0否1是（旧项目字段）',
                            `status` BOOLEAN DEFAULT NULL COMMENT '按钮权限状态(0无效1有效)',

                            `description` VARCHAR(255) DEFAULT NULL COMMENT '描述',
                            `create_by` VARCHAR(255) DEFAULT NULL COMMENT '创建人',
                            `update_by` VARCHAR(255) DEFAULT NULL COMMENT '更新人',
                            PRIMARY KEY (`id`),
                            KEY `index_menu_parent_id` (`parent_id`),
                            KEY `index_menu_type` (`menu_type`),
                            KEY `index_menu_hidden` (`hide_in_menu`),
                            KEY `index_menu_status` (`status`),
                            KEY `index_menu_path` (`path`),
                            KEY `index_menu_order_num` (`order_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单权限表(新结构)';

DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                  `parent_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '父权限ID',
                                  `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '权限名称',
                                  `uri` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '权限对应的资源URI或路径',
                                  `action` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '对资源执行的操作',
                                  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                  `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission` (
                                       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                       `role_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID，关联 sys_role 表',
                                       `permission_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID，关联 sys_permission 表',
                                       `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                       `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `uk_role_permission` (`role_id`, `permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色与权限的关联表';

DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `role_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID，关联 sys_role 表',
                                 `menu_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID，关联 sys_menu 表',
                                 `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                 `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `uk_role_menu` (`role_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色与菜单关联表';

DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict` (
                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                            `dict_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '字典名称',
                            `dict_code` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '字典编码',
                            `description` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
                            `is_deleted` BOOLEAN DEFAULT 0 COMMENT '是否删除，1删除0未删除',
                            `type` TINYINT UNSIGNED DEFAULT 0 COMMENT '字典类型0为string,1为number',
                            `create_by` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '创建人',
                            `update_by` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '更新人',
                            `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                            `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（为 NULL 表示未删除）',
                            PRIMARY KEY (`id`),
                            UNIQUE INDEX `uk_sd_dict_code` (`dict_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典表';

DROP TABLE IF EXISTS `sys_dict_item`;
CREATE TABLE `sys_dict_item` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `dict_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '字典ID，关联 sys_dict 表',
                                 `item_text` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '字典项文本',
                                 `item_value` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '字典项值',
                                 `description` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
                                 `sort_order` INT NOT NULL DEFAULT 0 COMMENT '排序',
                                 `is_enabled` BOOLEAN  DEFAULT 0 COMMENT '是否启用 1启用0不启用',
                                 `create_by` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '创建人',
                                 `update_by` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '更新人',
                                 `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                 `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（NULL表示未删除）',
                                 PRIMARY KEY (`id`),
                                 INDEX `idx_sdi_dict_id` (`dict_id`),
                                 INDEX `idx_sdi_sort_order` (`sort_order`),
                                 INDEX `idx_sdi_enabled` (`is_enabled`),
                                 INDEX `idx_sdi_dict_val` (`dict_id`, `item_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典数据表';
