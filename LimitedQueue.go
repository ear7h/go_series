package go_series

//LimitedQueue holds a slice which will always be the same length.
//It essentially acts like a shift register
//A sum variable is present to quickly retrieve averages
type LimitedQueue struct {
	q []float32
	sum float32
}

//Insert is a funtion for inserting data to the queue
//it returns the value which gets pushed out
func (l *LimitedQueue) Insert(n float32) float32 {

	o, s := l.q[0], l.q[1:]

	l.sum += n - o

	l.q = append(s, n)

	return o
}

//Avg returns the average of all the values int the queue
func (l *LimitedQueue) Avg() float32 {
	return l.sum / float32(len(l.q))
}

//NewLimitedQueue is a pseudo constructor
func NewLimitedQueue(s int) *LimitedQueue {
	r := &LimitedQueue{q: make([]float32, s), sum:0}
	return r
}
