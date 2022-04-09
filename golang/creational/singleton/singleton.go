package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton

// 双重检查的方式实现 singleton，这个方式叫做 Double Check Lock，也可以看成是
// Check-Lock-Check 的流程，这样的做法是想要尽可能地减少并发中竞争和同步的开销。
// 通过 race 竞态检测 `go test -race -bench=. -benchmem -run=none` 可知这种
// 实现方式存在数据竞争。
// var lock = &sync.Mutex{}

// func getInstance() *singleton {
// 	if instance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if instance == nil {
// 			fmt.Println("Creting Single Instance Now")
// 			instance = &singleton{}
// 		} else {
// 			fmt.Println("Single Instance already created-1")
// 		}
// 	} else {
// 		fmt.Println("Single Instance already created-2")
// 	}
// 	return instance
// }

// 通过 race 竞态检测 `go test -race -bench=. -benchmem -run=none` 可知这种
// 实现方式不存在数据竞争，所以该实现方式较 Double Check Lock 更优。
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		fmt.Println("Creting Single Instance Now")
		instance = &singleton{}
	})
	return instance
}
