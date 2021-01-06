// @Desc: 基于
// @CreateTime: 2021/1/6
package Week05

import "time"

// 存放队列
var LimitQueue map[string][]int64

func IsLimit(queueName string, limitNum int, limitDuring int64) bool {
    currTime := time.Now().Unix()
    // 如果为 nil 的话就进行初始化
    if LimitQueue == nil {
        LimitQueue = make(map[string][]int64, 0)
    }
    // 如果队列的值不存在的话就初始化
    if _, ok := LimitQueue[queueName]; !ok {
        LimitQueue[queueName] = make([]int64, 0)
    }
    // 如果队列未满的话就进行添加
    if len(LimitQueue[queueName]) < limitNum {
        LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
        return true
    }
    // 满了的话, 需要将第一个与当前进行比较, 如果大于设置的限制时间, 就需要剔除第一个值, 从新加入
    // 否则就不能
    earlyTime := LimitQueue[queueName][0]
    // 不在限定时间范围内
    if currTime - earlyTime > limitDuring {
        // 删除第一个
        LimitQueue[queueName] = LimitQueue[queueName][1:]
        LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
        return true
    }
    return false
}
