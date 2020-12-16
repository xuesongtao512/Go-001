// @Desc:
// @CreateTime: 2020/10/8
package user

import (
    "github.com/gin-gonic/gin"
    "gohomework.com/Week04/homework/app01/internal/controller"
    "gohomework.com/Week04/homework/app01/internal/modules/user"
    "gohomework.com/Week04/homework/app01/pkg/mylog"
    "net/http"
)


var log = mylog.Log

// login
func Login(c *gin.Context) {
    var userInfo user.User
    if err := c.BindJSON(&userInfo); err != nil {
        log.Error("c.BindJSON(userInfo) is failed, err: ", err)
        return
    }
    resp := controller.Response{
        Status: 1,
        Msg:    "登录成功",
    }
    if userInfo.Username == "test" {
        if userInfo.Password != "123456" {
            resp.Msg = "用户名或密码错误"
        }
    } else {
        resp.Msg = "用户名或密码错误"
    }
    c.JSON(http.StatusOK, &resp)
}
