// @Desc: errgroup 起 3 个 goroutine, 当一个报错时, 其他就停止
// @CreateTime: 2020/12/13
package main

import (
    "context"
    "errors"
    "golang.org/x/sync/errgroup"
    "log"
)

func main() {
    g, ctx := errgroup.WithContext(context.Background())
    var a int
    g.Go(func() error {
        a += 1
        return errors.New("my is failed")
    })

    g.Go(func() error {
        a += 2
        select {
        case <-ctx.Done():
            // 如果有报错的话就 - 2
            a -= 2

        }
        return nil
    })

    g.Go(func() error {
        a += 3
        select {
        case <-ctx.Done():
            // 如果有报错的话就 - 3
            a -= 3
        }
        return nil
    })

    if err := g.Wait(); err != nil {
        log.Println(err)
    }
    log.Println(a)
}

