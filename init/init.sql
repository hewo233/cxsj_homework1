-- 创建数据库 cxsj1db
CREATE DATABASE IF NOT EXISTS cxsj1db;

-- 使用数据库 cxsj1db
USE cxsj1db;

-- 创建表 users
CREATE TABLE IF NOT EXISTS users (
    id INT(11) NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) UNIQUE,
    name VARCHAR(100),
    password VARCHAR(255),
    gender CHAR(1),
    PRIMARY KEY (id)
);
