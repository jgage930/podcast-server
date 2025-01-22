package filestore

type Queue struct {
	store []Task
}

func NewQueue() Queue {
	queue := Queue{
		store: make([]Task, 0),
	}
	return queue
}

func (q *Queue) Push(task Task) {
	_ = append(q.store, task)
}

func (q *Queue) Pop() Task {
	task := q.store[0]
	q.store = q.store[1:]
	return task
}

func (q *Queue) IsEmpty() bool {
	return len(q.store) == 0
}
