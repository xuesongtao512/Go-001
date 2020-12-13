// @Desc: 起 3 个 goroutine, errgroup 等待所有 goroutine 执行结束, 不管报错与否
// @CreateTime: 2020/12/13
package main

import (
    "errors"
    "golang.org/x/sync/errgroup"
    "log"
)

func main() {
    g := new(errgroup.Group)
    var a int
    g.Go(func() error {
        a += 3
        return errors.New("my is failed")
    })

    g.Go(func() error {
        a += 1
        return nil
    })

    g.Go(func() error {
        // log.Println("my is ok")
        a += 2
        return errors.New("xxx")
    })

    if err := g.Wait(); err != nil {
        log.Println(err)
    }
    log.Println(a)
}

