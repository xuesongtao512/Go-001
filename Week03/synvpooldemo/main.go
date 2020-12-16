// @Desc:
// @CreateTime: 2020/12/13
package main

import (
    "log"
    "sync"
)

func main() {
    pipe := &sync.Pool{New: func() interface{} {
        return "hello"
    }}

    val := "h"

    pipe.Put(val)

    log.Println(pipe.Get())
    log.Println(pipe.Get())
    log.Println(pipe.Get())
    log.Println(pipe.Get())
}