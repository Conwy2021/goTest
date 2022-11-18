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

/*
Go语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。
这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。
*/
func Test202210280930(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	//ch <- 3 //会报错，因为缓冲区只允许大小为2
	x1 := <-ch

	fmt.Println(x1)

}
