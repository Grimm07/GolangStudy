package Queue

import "go/types"

type Queue struct {
	cap   int
	count int
	queue []types.Object
}

func (q Queue) initialize(qType types.Object, qCap int) (ret Queue) {
	ret = q
	return new(Queue{cap: qCap, count: 0, var queue = [qCap]qType})
}

func (q Queue) enqueue(itm types.Object) {
	q.queue[q.count] = itm
	q.count++
}

func (q Queue) dequeue() (x types.Object) {
	x = q.queue[0]
	return
}
