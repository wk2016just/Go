package main
import "fmt"
import "time"
// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. 
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)  //这里模拟任务就是sleep
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2   //results存储任务执行结果
    }
}

func main() {

    // In order to use our pool of workers we need to send
// them work and collect their results. 
//分别是为了发送任务的jobs参数通道，和存储执行结果的results通道
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // This starts up 3 workers, initially blocked
    // because there are no jobs yet. 初始被2个通道阻塞
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Here we send 5 `jobs` and then `close` that
    // channel to indicate that's all the work we have.
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Finally we collect all the results of the work.
    for a := 1; a <= 5; a++ {
        <-results
    }}
    
    