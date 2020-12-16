// @Desc:
// @CreateTime: 2020/12/16
package controller

// 响应体
type Response struct {
    Status uint `json:"status,omitempty"`   // 0 - 失败, 1 - 成功
    Msg string `json:"msg,omitempty"`
    Data interface{} `json:"data,omitempty"`
    Total int `json:"total,omitempty"` // 总条数
    CurrentPage int `json:"current_page,omitempty"` // 当前页
    TotalPage int `json:"total_page,omitempty"` // 总页数
}
