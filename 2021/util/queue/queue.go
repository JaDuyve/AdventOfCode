package queue

type OrderedQueue struct {
	Queue []int
	Front int
	Back  int
}

func NewQueue(capacity int) *OrderedQueue {
	queue := new(OrderedQueue)
	queue.Queue = make([]int, capacity)

	return queue
}

func (queue *OrderedQueue) Add(value int) {
	if queue.Back == len(queue.Queue) {
		newQueue := make([]int, 2*len(queue.Queue))
		copy(newQueue, queue.Queue)
		queue.Queue = newQueue
	}

	for i, x := range queue.Queue {
		if x < value {
			copy(queue.Queue[i+1:], queue.Queue[i:])
			queue.Queue[i] = value
			break
		}
	}

	queue.Back++
}

func (queue *OrderedQueue) Remove() int {
	if queue.Size() == 0 {
		panic("No such element")
	}

	data := queue.Queue[queue.Front]
	queue.Queue[queue.Front] = 0
	queue.Front++
	if queue.Size() == 0 {
		queue.Front = 0
		queue.Back = 0
	}
	return data
}

func (queue *OrderedQueue) Peek() int {
	if queue.Size() == 0 {
		panic("No such element")
	}

	return queue.Queue[queue.Front]
}

func (queue *OrderedQueue) Size() int {
	return queue.Back - queue.Front
}
