package main

import (
	"fmt"
	"golang-package/channel"
	"golang-package/unique_id_factory"
	"time"
)

func main() {
	queueInit()
}

func queueInit() {
	queue := channel.NewQueue(10)

	listenMethod := func(val int, workerNum int) {
		fmt.Printf("workNum: %d val: %d\n", workerNum, val)
	}

	queue.Listen(1000 * time.Millisecond, listenMethod, 1)
	queue.Listen(1000 * time.Millisecond, listenMethod, 2)
	queue.Listen(1000 * time.Millisecond, listenMethod, 3)

	idx := 0
	for {
		idx++
		data := channel.NewData()
		val, err := unique_id_factory.GetID()

		if err != nil {
			continue
		}

		data.Val = int(val)
		queue.Push(data)
		fmt.Printf("queue len: %d\n", queue.Length())
		time.Sleep(300 * time.Millisecond)
	}
}
