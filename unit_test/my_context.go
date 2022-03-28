package unit_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func httpRequest(ctx context.Context) {
	for {
		// process http requests
		select {
		case <-ctx.Done():
			fmt.Println("http requests cancel")
			return
		case <-time.After(time.Second * 1):
			fmt.Println("http requests 1")
		}
	}

}

func TestTimeoutContext(t *testing.T) {
	fmt.Println("start TestTimeoutContext")
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	//defer cancel()
	//httpRequest(ctx)
	//fmt.Println("http requests 2")
	//time.Sleep(time.Second * 5)

	var wg sync.WaitGroup
	messages := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//defer close(messages)
	defer cancel()

	// consumer
	wg.Add(1)
	go func(ctx context.Context) {
		// 每隔1s 打印一次message, 直到第5s timeout, 中断所有的goroutine.
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)

	//select {
	//case <-ctx.Done():
	//	fmt.Println("main process exit!")
	//}
	wg.Wait()
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case ch <- n:
				n++
				fmt.Println("n:", n)
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println("canceled")
				return
			}
		}
	}()
	return ch
}

func TestCancelContext(t *testing.T) {
	// 创建一个Cancel context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			// 达到要求后触发 cancel
			cancel()
			break
		}
	}
}

func TestCancelContextV2(t *testing.T)  {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func1(ctx, wg)
	time.Sleep(time.Second * 2)
	// 人为触发取消
	cancel()
	// 等待goroutine退出
	wg.Wait()
}

func func1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	respC := make(chan int)
	defer close(respC)
	// 处理逻辑 无缓冲channel 需要先读取，再写入，不然会死锁，所以加到goroutine里写，先阻塞挂起，让后面的读取先执行.
	go func() {
		time.Sleep(time.Second * 1)  // 模拟耗时1s后，再写入channel
		respC <- 10
	}()
	// 取消机制
	for  {
		select {
		case <-ctx.Done():  // 当c.done被关闭后，Done()返回{}，取消阻塞，return func, 结束所有goroutine.
			fmt.Println("child process cancel")
			return
		case r := <-respC:
			fmt.Println(r)
		}
	}
}

