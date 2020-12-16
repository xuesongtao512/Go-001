// @Desc:
// @CreateTime: 2020/12/15
package main

import (
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "context"

    "github.com/gin-gonic/gin"
    "github.com/fvbock/endless"

    config "gohomework.com/Week04/homework/app01/configs"
    router2 "gohomework.com/Week04/homework/app01/internal/rotuer"
    "gohomework.com/Week04/homework/app01/pkg/conn"
    log2 "gohomework.com/Week04/homework/app01/pkg/mylog"

)

func main() {
    // 初始化 mylog
    log2.InitLogger()
    // 初始化配置信息
    projectDir, err := os.Getwd()
    if err != nil {
        log2.Log.Fatal("get projectDir is failed, err: ", err)
    }
    config.InitConf(projectDir)
    // 初始化数据库
    if err := conn.InitDB(); err != nil {
        log2.Log.Fatal("connect DB failed, err:%v\n", err)
    }

    router := gin.Default()
    router2.Routers(router)

    // 监听关机
    signalCh := make(chan os.Signal)
    signal.Notify(signalCh, syscall.SIGINT | syscall.SIGTERM)
    if err := servers(signalCh, ":9999", router); err != nil {
        log2.Log.Fatal("listen server is failed, err: ", err)
    }
}

func servers(signalCh chan os.Signal, host string, router *gin.Engine) error {
    server := &http.Server{
        Addr:              host,
        Handler:           router,
    }
    go func() {
        select {
        case <-signalCh:
            server.Shutdown(context.Background())
        }
    }()
    return server.ListenAndServe()
}

