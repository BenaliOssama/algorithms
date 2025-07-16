package queue

// define a type
type Node struct {
	Item any
	Prev *Node
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

// return pointer or actual variable ?
func New() *Queue { // {addr 1001}// head
	return &Queue{Head: nil, Tail: nil, Size: 0}
}

func (q *Queue) Dequeue() *Node {
	if q.Size == 0 {
		return nil
	}
	dq := q.Head
	if q.Head.Next == nil {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = dq.Next
		q.Head.Prev = nil
	}
	q.Size--

	return dq
}

func (q *Queue) Enqueue(item any) {
	node := &Node{Item: item, Prev: nil, Next: nil}
	q.Size++
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		node.Prev = q.Tail
		q.Tail = node
	}
}

func (q *Queue) Peek() any {
	if q.Tail == nil {
		return nil
	}
	return q.Tail.Item
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
