// @Desc:
// @CreateTime: 2020/12/13
package main

import (
    "context"
    "fmt"
    "time"
)

const shortDuration = time.Microsecond

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
    defer cancel()

    select {
    case <-time.After(time.Second):
        fmt.Println("over")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }

}
