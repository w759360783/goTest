package spiderkit

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randomMT sync.Mutex
)
/**
生成start-end之间的随机数
 */

func GetRandomInt(start, end int) int {
	randomMT.Lock()
	defer randomMT.Unlock()
	// 延迟毫秒防止使用重复的随机数种子
	<-time.After(time.Nanosecond)
	i := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + i.Intn(end-start)
}
