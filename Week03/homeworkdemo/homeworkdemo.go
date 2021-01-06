// @Desc:
// @CreateTime: 2020/12/19
package homeworkdemo

import (
    "context"
)

type Hook struct {
     OnStart func(ctx context.Context) error
     OnStop func(ctx context.Context) error
}

type App struct {
    hooks []Hook
}

func NewApp() *App {
    return &App{}
}

func (a *App) name()  {
    
}
