package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")
	elapsed := time.Since(start)
	fmt.Println("耗时：", elapsed)
}
