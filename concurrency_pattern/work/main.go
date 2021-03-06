// 代码来自《GO语言实战》
// 这个示例程序展示如何使用work包创建一个goroutine池并完成工作
package main

import (
	"./work"
	"log"
	"sync"
	"time"
)

// names提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个goroutine来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{name: name}
			go func() {
				// 将任务提交执行，当Run返回我们就知道任务已经处理完了
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
