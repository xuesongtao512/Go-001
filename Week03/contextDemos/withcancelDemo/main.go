// @Desc:
// @CreateTime: 2020/12/13
package main

import (
    "context"
    "fmt"
)

func main() {
    gen := func(ctx context.Context) <- chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return
                case dst <- n:
                    n ++
                }
            }
        }()
        return dst
    }

    ctx, cancle := context.WithCancel(context.Background())
    defer cancle()

    for n:= range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
