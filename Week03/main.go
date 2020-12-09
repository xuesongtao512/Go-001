// @Desc:
// @CreateTime: 2020/12/9
package main

import (
    "context"
    "fmt"
    "github.com/pkg/errors"
    "golang.org/x/sync/errgroup"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func router() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
            log.Println("recv is failed, err: ", err)
        }
        fmt.Fprintln(w, string(body))
    })
    return mux
}

// 定义一个 server
func server(ctx context.Context, signal  <-chan os.Signal, addr string, mux *http.ServeMux) error {
    server := http.Server{
        Addr:              addr,
        Handler:           mux,
    }

    go func() {
        select {
        case <-ctx.Done():
            log.Println("cancel server!")
        case <-signal:
            log.Println("signal is stop!")
        }
        server.Shutdown(ctx)
    }()
    log.Println("server is running...")
    return server.ListenAndServe()
}

func main() {
    r := router()
    signalCh := make(chan os.Signal)
    // TODO 只监听了两种
    signal.Notify(signalCh, syscall.SIGKILL | syscall.SIGTERM)
    g, ctx := errgroup.WithContext(context.Background())
    g.Go(func() error {
         if err := server(ctx, signalCh, ":8080", r); err != nil {
             return errors.Wrap(err, "server run is failed")
         }
         return nil
    })
    if err := g.Wait(); err != nil {
        log.Printf("server run failed, err: %#v", err)
    }
}
