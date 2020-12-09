// @Desc:
// @CreateTime: 2020/11/29
package dao

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/pkg/errors"
    "log"
)

var DB *sqlx.DB


// init db
func InitDb() error {
    var err error
    DB, err = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
    if err != nil {
        return errors.Wrap(err, "Connect db")
    }
    DB.SetMaxOpenConns(100) // 设置连接池最大连接数
    DB.SetMaxIdleConns(20)  // 设置连接池最大空闲连接数

    log.Println("init DB is ok...")
    return nil
}

// select user id
func SelectUserId() (int, error) {
    var userId int
    err := DB.Get(&userId, `SELECT id FORM t_user WHERE name = ?`, "test")
    if err != nil {
        return 0, errors.Wrapf(err, `SELECT id FORM t_user WHERE name = ?`, "test")
    }
    return userId, nil
}
