// @Desc:
// @CreateTime: 2020/12/16
package conn

import (
    "github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() (err error) {
    dsn := "user:password@tcp(127.0.0.1:3306)/test"
    // 也可以使用MustConnect连接不成功就panic
    Db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        return err
    }
    // 设置最大连接数
    Db.SetMaxOpenConns(20)
    // 设置最大闲置数
    Db.SetMaxIdleConns(10)
    return
}