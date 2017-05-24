package go_series

type LimitedQueue struct {
	q []float32
	sum float32
}

func (l *LimitedQueue) Insert(n float32) float32 {

	o, s := l.q[0], l.q[1:]

	l.sum += n - o

	l.q = append(s, n)

	return o
}

func (l *LimitedQueue) Avg() float32 {
	return l.sum / float32(len(l.q))
}

func NewLimitedQueue(s int) *LimitedQueue {
	r := &LimitedQueue{q: make([]float32, s), sum:0}
	return r
}

func makeLimitedQueue(s int) LimitedQueue {
	r := LimitedQueue{q: make([]float32, s), sum:0}
	return r
}
