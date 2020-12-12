// @Desc:
// @CreateTime: 2020/12/12
package testInternal

import (
    "gohomework.com/Week04/studyDemo/server"
    "testing"
    "gohomework.com/Week04/studyDemo/internal/utils"
)

// 导入 studyDemo/internal/utils.go
func TestImportUtils(t *testing.T)  {
    // utils.IsZero(0)
}

func TestImportServer(t *testing.T)  {
    server.MyServer()
}
