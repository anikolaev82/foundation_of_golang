package stopwatch

import "time"

type StopWatch struct {
	start   time.Time
	stop    time.Time
	results []time.Duration
}

func (sw *StopWatch) Start() {
	sw.start = time.Now()
	sw.results = nil
}

func (sw *StopWatch) Stop() {
	sw.stop = time.Now()
}

func (sw *StopWatch) SaveSplit() {
	sw.results = append(sw.results, time.Now().Sub(sw.start))
}

func (sw *StopWatch) GetResults() []time.Duration {
	return sw.results
}
