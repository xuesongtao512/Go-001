// @Desc:
// @CreateTime: 2020/10/8
package router

import (
    "github.com/gin-gonic/gin"
    "gohomework.com/Week04/homework/app01/internal/controller/user"
)


func Routers(router *gin.Engine) {
    router.POST("/api/login", user.Login)
}

