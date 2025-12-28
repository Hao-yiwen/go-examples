-- Go-Zero 微服务用户管理系统 数据库初始化脚本

-- 创建数据库
CREATE DATABASE IF NOT EXISTS go_zero_user DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE go_zero_user;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    username VARCHAR(50) NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码(加密)',
    email VARCHAR(100) DEFAULT '' COMMENT '邮箱',
    phone VARCHAR(20) DEFAULT '' COMMENT '手机号',
    avatar VARCHAR(255) DEFAULT '' COMMENT '头像URL',
    status TINYINT UNSIGNED DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY uk_username (username),
    KEY idx_status (status),
    KEY idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 角色表
CREATE TABLE IF NOT EXISTS roles (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '角色ID',
    name VARCHAR(50) NOT NULL COMMENT '角色名称',
    code VARCHAR(50) NOT NULL COMMENT '角色编码',
    description VARCHAR(255) DEFAULT '' COMMENT '角色描述',
    status TINYINT UNSIGNED DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY uk_code (code),
    KEY idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS user_roles (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    role_id BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY uk_user_role (user_id, role_id),
    KEY idx_user_id (user_id),
    KEY idx_role_id (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 初始化管理员角色
INSERT INTO roles (name, code, description) VALUES
    ('管理员', 'admin', '系统管理员，拥有所有权限'),
    ('普通用户', 'user', '普通用户，拥有基本权限')
ON DUPLICATE KEY UPDATE name = VALUES(name);

-- 初始化管理员用户 (密码: admin123, 使用bcrypt加密)
-- 实际密码需要在程序中加密后插入
INSERT INTO users (username, password, email, status) VALUES
    ('admin', '$2a$10$EqKcp1WaEHPejGWbhF.8dOJwJYpTVrQHmXAWWYcGC8xN.TgxH5Npe', 'admin@example.com', 1)
ON DUPLICATE KEY UPDATE email = VALUES(email);

-- 为管理员分配管理员角色
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u, roles r WHERE u.username = 'admin' AND r.code = 'admin'
ON DUPLICATE KEY UPDATE user_id = user_id;
