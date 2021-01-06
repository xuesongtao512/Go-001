// @Desc:
// @CreateTime: 2021/1/6
package Week05

import (
    "log"
    "testing"
)

func TestIsLimit(t *testing.T) {
    for i := 0; i < 15; i++ {
        isOk := IsLimit("test", 10, 5)
        if isOk {
            log.Println("运行正常")
        } else {
            log.Println("运行限制")
        }
    }
}