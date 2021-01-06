// @Desc: 基于时间滑动窗口实现
// @CreateTime: 2021/1/4
package Week05

import "time"

type RateLimit struct {
	appName     string  // 应用名
	limitQueue  []int64 // 存放队列
	limitNum    int     // 限制数量
	limitDuring int64   // 限制的时间
}

// 初始化对象
func NewRateLimit(appName string, limitNum int, limitDuring int64) *RateLimit {
	return &RateLimit{
		appName:     appName,
		limitQueue:  make([]int64, 0),
		limitNum:    limitNum,
		limitDuring: limitDuring,
	}
}

// 获取令牌
func (r *RateLimit) GetToken() bool {
	currTime := time.Now().Unix()
	// 如果队列未满的话就进行添加
	if len(r.limitQueue) < r.limitNum {
		r.limitQueue = append(r.limitQueue, currTime)
		return true
	}
	// 满了的话, 需要将第一个与当前进行比较, 如果大于设置的限制时间, 就需要剔除第一个值, 从新加入
	// 否则就不能
	earlyTime := r.limitQueue[0]
	// 不在限定时间范围内
	if currTime-earlyTime > r.limitDuring {
		// 删除第一个
		r.limitQueue = r.limitQueue[1:]
		r.limitQueue = append(r.limitQueue, currTime)
		return true
	}
	return false
}
