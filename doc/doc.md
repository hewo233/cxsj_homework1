# 项目文档

## 项目介绍
这是一个由 golang 编写的简单的用户信息管理系统，使用了 mariadb 作为数据库。
(有点放飞自我了，密码加密隐藏什么的后面写（)

## 项目结构
```bash
.
├── common
│   └── database.go
├── doc
│   └── doc.md
├── go.mod
├── go.sum
├── main.go
├── model
│   └── dbModel.go
└── README.md

4 directories, 7 files
````

其中，common 文件夹存放了数据库操作的公共函数，model 文件夹存放了数据库模型。
由于是一个初始项目，所以操作放置于 main.go 中。

## API

```go
func InitDB(dbPassword string) (*sql.DB, error)
func CloseDB(db *sql.DB) error
func AddRecord(info *model.CxsjUser, db *sql.DB) error
func QueryRecordByEmail(email string, db *sql.DB) (*model.CxsjUser, error)
func ModifyRecord(info *model.CxsjUser, db *sql.DB)
func DeleteRecord(email string, db *sql.DB) error
```

## 数据库模型

```mysql
MariaDB [cxsj1db]> DESCRIBE users;
+----------+--------------+------+-----+---------+----------------+
| Field    | Type         | Null | Key | Default | Extra          |
+----------+--------------+------+-----+---------+----------------+
| id       | int(11)      | NO   | PRI | NULL    | auto_increment |
| email    | varchar(100) | YES  | UNI | NULL    |                |
| name     | varchar(100) | YES  |     | NULL    |                |
| password | varchar(255) | YES  |     | NULL    |                |
| gender   | char(1)      | YES  |     | NULL    |                |
+----------+--------------+------+-----+---------+----------------+

```

## 使用方法

本地环境下， 运行后输入数据库密码，即可使用。
操作可见 main.go 。