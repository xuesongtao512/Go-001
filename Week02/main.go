// @Desc:
// @CreateTime: 2020/11/29
package main

import (
    "log"
    "my_go/Week02/dao"
)

func main() {
    if err := dao.InitDb(); err != nil {
        log.Fatalf("init DB is failed, err: %#v", err)
    }
}
