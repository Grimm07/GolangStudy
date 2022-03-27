/*
	IT 327 Project - Golang Study
	Queue implementation (of type integer, but can be anything)
	Testing occurs in Test
*/

package Queue

const MAX_SIZE = 1000

type Queue[T any] struct {
	count int
	Q     [MAX_SIZE]T
}

func (q *Queue[T]) Enqueue(itm T) {
	q.Q[q.count] = itm
	q.count++
}

func (q *Queue[T]) Dequeue() (x T) {
	x = q.Q[0]
	for i := 0; i < q.count; i++ {
		q.Q[i] = q.Q[i+1]
	}
	q.count--
	return
}

func (q *Queue[T]) IsEmpty() bool { return q.count == 0 }

func (q *Queue[T]) Peek() T { return q.Q[0] }

func (q *Queue[T]) IsFull() bool { return q.count == MAX_SIZE-1 }

func (q *Queue[T]) Size() int { return q.count }
