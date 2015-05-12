package mapreduce

import (
	"testing"
)

var mr *MapReducer

func (mr *MapReducer) Map(in chan interface{}, out chan interface{}) {
	for v := range in {
		out <- v
	}
	close(out)
}

func (mr *MapReducer) Reduce(in chan interface{}) interface{} {
	res := 0
	for v := range in {
		res += v.(int)
	}
	return res
}

func TestMapReduce(t *testing.T) {
	inChan := make(chan interface{})
	outChan := make(chan interface{})

	config := Configuration{
		mapperCount: 1,
		inChan:      inChan,
		outChan:     outChan,
	}

	mr = newMapReducer(config)

	// Feed input channel
	go func(in chan interface{}) {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}(inChan)

	result, _ := mr.Run()

	if result != 45 {
		t.Fail()
	}
}
