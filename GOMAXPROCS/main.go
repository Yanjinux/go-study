package main

import (
	"fmt"
	"runtime"
	"sync"
)

const count = 258 // 258

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(count)
	/*
		P {
			runnext : 257
		    runq : 1 2 . . . . 256
		}

	*/

	/*
		258 时   p.schedtick%61 就会去 全局p取 因此 在执行local q时会穿插部分G
	*/
	for i := 1; i <= count; i++ {
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
