package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// start := time.Now()
	// timer1 := time.NewTimer(5 * time.Second)
	// <-timer1.C
	// fmt.Println("時間です")
	// end := time.Now()
	// fmt.Println("%f秒", (end.Sub(start)).Seconds())
	ch := make(chan string)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("開始")
	go channelFunction(ch)

	fmt.Println("処理中")
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(<-ch)

	fmt.Println("終了しました")
	fmt.Println(runtime.NumGoroutine())
}

func channelFunction(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "done"
}
