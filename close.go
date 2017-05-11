package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine.  In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d
	go func() {
		for {
			j, more := <-jobs //这里是双赋值初始化，j,more都接收jobs的值
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.(done阻塞在这里主要是为了防止worker还没有接收完成主线程就结束）
	<-done
}
