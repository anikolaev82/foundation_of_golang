package circularbuffer

import "fmt"

type TypeValue string

var DefVal TypeValue

type CircularBuffer struct {
	values  []TypeValue
	headIdx int
	tailIdx int
}

func (cb *CircularBuffer) Push(val TypeValue) {
	cb.values[cb.tailIdx] = val
	cb.tailIdx = (cb.tailIdx + 1) % cap(cb.values)
	if cb.tailIdx == cb.headIdx {
		cb.headIdx = (cb.headIdx + 1) % cap(cb.values)
	}
}

func (cb *CircularBuffer) GetValues() (ret []TypeValue) {
	fmt.Printf("head %d tail %d cap %d\n", cb.headIdx, cb.tailIdx, cap(cb.values))

	for item := cb.headIdx; item != cb.tailIdx; item = (item + 1) % cap(cb.values) {
		ret = append(ret, cb.values[item])
	}

	return
}

func (cb *CircularBuffer) Pop() (ret TypeValue) {
	ret = cb.values[cb.headIdx]
	cb.values[cb.headIdx] = DefVal
	cb.headIdx = (cb.headIdx + 1) % cap(cb.values)

	return
}

func NewCircularBuffer(capacity int) (cb *CircularBuffer) {
	capacity += 1
	cb = &CircularBuffer{values: make([]TypeValue, capacity, capacity)}

	return
}
