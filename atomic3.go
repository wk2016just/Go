package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single  这个例子中状态被单独的线程所拥有
// goroutine. This will guarantee that the data is never 确保数据不会被同时读写导致错误
// corrupted with concurrent access. In order to read or
// write that state, other goroutines will send messages
// to the owning goroutine and receive corresponding  其他线程给拥有状态的线程发送信息接收其回应
// replies. These `readOp` and `writeOp` `struct`s read
// encapsulate those requests and a way for the owning
// goroutine to respond.		 readOp和wirteOP封装了这个请求
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// As before we'll count how many operations we perform. 计算做了多少次操作
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// The `reads` and `writes` channels will be used by
	// other goroutines to issue read and write requests,
	// respectively.    reads和writes通道被其他go线程使用发布各自的请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// Here is the goroutine that owns the `state`, which
	// is a map as in the previous example but now private  goroutine拥有state，现在变成了goroutine私有的map结构
	// to the stateful goroutine. This goroutine repeatedly
	// selects on the `reads` and `writes` channels,  这个goroutine反复选择reads和writes通道回应请求
	// responding to requests as they arrive. A response
	// is executed by first performing the requested  一个响应被执行 被这个第一个
	// operation and then sending a value on the response
	// channel `resp` to indicate success (and the desired
	// value in the case of `reads`).
	go func() { //这里的线程和外界唯一联系就是reads和writes通道，reads读取由线程拥有的状态，writes去写他;
		//这里利用for select实现不断的监听读取同时也保证了要么读要么写其中一种
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads: //reads是*readOp通道，接受readOp类型值，然后把状态的int值赋给read的resp
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val //把val写入当前状态，然后把write.resp即chan bool置为true
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the
	// state-owning goroutine via the `reads` channel.
	// Each read requires constructing a `readOp`, sending
	// it over the `reads` channel, and the receiving the
	// result over the provided `resp` channel. 用reads发送读请求，用resp接收结果
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				fmt.Println(<-read.resp)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar
	// approach.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.
	time.Sleep(time.Second)

	// Finally, capture and report the op counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
