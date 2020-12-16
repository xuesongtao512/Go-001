// @Desc:
// @CreateTime: 2020/12/13
package main

import (
    "context"
    "fmt"
)

func setUserName() context.Context {
    ctx := context.WithValue(context.Background(), "Username", "test")
    return ctx
}

func main() {
    ctx := setUserName()
    fmt.Println(ctx.Value("Username"))
}
