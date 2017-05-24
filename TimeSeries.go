package go_series


//TimeSeries this struct is an aggregation of multiple limited queues.
//the individual queues represent a period of time which would like to be observable.
//given the Avg() funtion of LimitedQueue time series is capable of storing
//moving averages over a specified period of time for some number of periods
//sum is the sum of all items in the series.
//i is the interval or period, the total number of items to consider in the siliding average
//l is the length or total number of periods to observe.
type TimeSeries struct {
	q []*LimitedQueue //source
	sum float32
	i int //interval
	l int //length
}

//Insert works similarly to the function in LimitedQueue but due to the
//structure of TimeSeries, the LimitedQueues must be placed ass to mouth as in human centipede.
//the output of this function is the value which gets shifted out
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

//Avg returns the average all values in the time series
func (t *TimeSeries) Avg() float32 {

	return t.sum/ float32(t.i * t.l)
}

//SectionAvg returns the average of the LimitedQueue specified by s.
//it is important to note that the LimitedQueues stored in the TimeSeries are
//in chronological. This is different than a single limited queue where newer values
//are appended to underlying slice
func (t *TimeSeries) SectionAvg(s int) float32 {
	return t.q[s].Avg()
}

//Section returns a []float32 with the values of section specified by s.
//it is important to note that the LimitedQueues stored in the TimeSeries are
//in chronological. This is different than a single limited queue where newer values
//are appended to underlying slice
func (t *TimeSeries) Section(s int) []float32 {
	return t.q[s].q
}

//NewTimeSeries is the constructor for TimeSeries, it requires the period and length
//for the time series
func NewTimeSeries(i, l int) *TimeSeries {
	arr := make([]*LimitedQueue, l)
	for k := range arr{
		arr[k] = NewLimitedQueue(i)
	}
	r := &TimeSeries{q:arr ,i:i, l:l}

	return r
}
