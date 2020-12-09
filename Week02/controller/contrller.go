// @Desc:
// @CreateTime: 2020/11/30
package controller

import (
    "github.com/pkg/errors"
    "log"
    "my_go/Week02/server"
)

type Response struct {
    Code int `json:"code,omitempty"` // 0 - fail, 1 - success
    Msg string `json:"msg,omitempty"`
    Data interface{} `json:"data,omitempty"`
}


func GetUserId()  {
    userId, err := server.GetUserId();
    resp := &Response{Code: 1}
    if err != nil {
        resp.Code = 0
        resp.Msg = "您查询 user_id 不存在"
        log.Println("select userId is failed, err: %v", errors.Cause(err))
    } else {
        resp.Data = map[string]int{"user_id": userId}
    }
    // TODO json.Marshal() 后并响应
}