package data_structures

import "fmt"

type Queue struct {
	data 	[]int
	first 	int
	last		int
	size		int
	cap			int
}

func MakeQueue(size int) (*Queue, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size invalid")
	}
	return &Queue{make([]int, size), -1, -1, 0, size}, nil
}

func (queue *Queue) Enqueue(val int) error {
	if queue.size == queue.cap {
		return fmt.Errorf("queue is full")
	}

	queue.last = (queue.last + 1) % queue.cap
	queue.data[queue.last] = val
	queue.size++

	if queue.first < 0 { // empty queue at first
		queue.first = queue.last
	}

	return nil;
}

func (queue *Queue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return -1, fmt.Errorf("empty queue")
	}

	val := queue.data[queue.first]

	queue.data[queue.first] = -1
	queue.size--
	if queue.size == 0 {
		queue.first = -1
		queue.last = -1
	} else {
		queue.first++
	}

	return val, nil
}

func (queue Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue Queue) PrintQueue() {
	fmt.Println(queue.data);
	if (queue.size > 0) {
		fmt.Println("First:", queue.data[queue.first])
		fmt.Println("Last:", queue.data[queue.last])
	}
}