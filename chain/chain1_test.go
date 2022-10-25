package chain_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test1013(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 //会报错，因为缓冲区只允许大小为2
	x1 := <-ch
	x2 := <-ch
	fmt.Println(x1)
	fmt.Println(x2)
}

func Test10210933(t *testing.T) {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}

func Test10210935(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}
