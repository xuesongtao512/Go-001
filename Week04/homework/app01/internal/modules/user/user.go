// @Desc:
// @CreateTime: 2020/10/11
package user

// user 模型
type User struct {
     Name string `json:"name"` // 用户名
     Username string `json:"username"`
     Password string `json:"password"`
     Phone string `json:"phone,omitempty"`
     Addr string `json:"addr,omitempty"`
}

