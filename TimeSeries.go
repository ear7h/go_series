package go_series



type TimeSeries struct {
	q []*LimitedQueue //source
	sum float32
	i int //interval
	l int //length
}

func (t *TimeSeries) Insert(f float32) float32 {
	i := f //insert value
	sum := float32(0) //sum for time series
	for _, v := range t.q {
		i = v.Insert(i)
		sum += v.sum
	}
	t.sum = sum
	return i
}

func (t *TimeSeries) Avg() float32 {

	return t.sum/ float32(t.i * t.l)
}

func (t *TimeSeries) SectionAvg(s int) float32 {
	return t.q[s].Avg()
}

func (t *TimeSeries) Section(s int) []float32 {
	return t.q[s].q
}
func NewTimeSeries(i, l int) *TimeSeries {
	arr := make([]*LimitedQueue, l)
	for k := range arr{
		arr[k] = NewLimitedQueue(i)
	}
	r := &TimeSeries{q:arr ,i:i, l:l}

	return r
}
