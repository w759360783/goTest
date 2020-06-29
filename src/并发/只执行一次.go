package main

import (
	"fmt"
	"sync"
	"time"
)

type Man struct {
	Name string
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	m := &Man{"比尔"}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				Kill(m)
			})
			wg.Done()
		}()
	}
	wg.Wait()
}
func Kill(m *Man) {
	<-time.After(3 * time.Second)
	fmt.Println(m.Name, ":啊啊啊啊啊，我死了呀")
}
