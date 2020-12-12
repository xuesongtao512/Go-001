// @Desc:
// @CreateTime: 2020/12/9
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "golang.org/x/sync/errgroup"
)

type myServer struct {
    app  *http.ServeMux
    host string
}

func app1() *http.ServeMux {
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

func app2() *http.ServeMux {
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
func server(ctx context.Context, signal <-chan os.Signal, addr string, mux *http.ServeMux) error {
    server := http.Server{
        Addr:    addr,
        Handler: mux,
    }

    go func() {
        select {
        case <-ctx.Done():
            log.Println("cancel server!")
        case <-signal:
            log.Println("signal is stop!")
        }
        server.Shutdown(context.TODO())
    }()
    log.Println("server is running...")
    return server.ListenAndServe()
}

func main() {
    a1 := &myServer{app: app1(), host: ":8080"}
    a2 := &myServer{app: app2(), host: "8081"}
    registerServer := make([]*myServer, 0, 2)
    registerServer = append(registerServer, a1, a2)
    signalCh := make(chan os.Signal)
    signal.Notify(signalCh, syscall.SIGKILL|syscall.SIGTERM)
    g, ctx := errgroup.WithContext(context.Background())
    for _, s := range registerServer {
        g.Go(func() error {
            return server(ctx, signalCh, s.host, s.app)
        })
    }

    if err := g.Wait(); err != nil {
        log.Printf("server run failed, err: %#v", err)

    }
}
