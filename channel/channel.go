package channel

import "time"

type data struct {
	Val int
}

type queue struct {
	c chan *data
}

type ListenMethod func(val int, workerNum int)

// 0 equal new unbuffered chan
// NOTICE: unbuffered "寫入"需等待"被讀取"才算寫入完成
func NewQueue(bufferSize int) *queue {
	if bufferSize == 0 {
		return &queue{
			c: make(chan *data),
		}
	}

	return &queue{
		c: make(chan *data, bufferSize),
	}
}

func NewData() *data {
	return &data{}
}

func (queue *queue) Push(d *data) *queue {
	queue.c <- d
	return queue
}

func (queue *queue) Pop() *data {
	return <-queue.c
}

func (queue *queue) Listen(sleepInterval time.Duration, method ListenMethod, workerNum int) {
	go func() {
		for {
			d := queue.Pop()
			go method(d.Val, workerNum)
			time.Sleep(sleepInterval)
		}
	}()
}

func (queue *queue) Length() int {
	return len(queue.c)
}
