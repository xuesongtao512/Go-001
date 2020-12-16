// @Desc:
// @CreateTime: 2020/12/13
package main

import (
    "context"
    "fmt"
    "time"
)

const shortDuration = 1 * time.Microsecond

func main() {
    d := time.Now().Add(shortDuration)
    ctx, cancel := context.WithDeadline(context.Background(), d)

    defer cancel()
    select {
    case <-time.After(1*time.Second):
        fmt.Println("over")
    case <-ctx.Done():
        fmt.Println(ctx.Err())

    }
}
