// @Desc:
// @CreateTime: 2021/1/6
package Week05

import (
	"log"
	"testing"
)

func TestRateLimit_GetToken(t *testing.T) {
	// 初始化一个 10s 内只有 10 个的限流器
	rateLimit := NewRateLimit("test", 10, 10)
	for i := 0; i < 15; i++ {
		isOk := rateLimit.GetToken()
		if isOk {
			log.Println("运行正常")
		} else {
			log.Println("运行限制")
		}
	}
}
