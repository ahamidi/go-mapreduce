package mapreduce

import (
	"testing"
)

type MapReducer struct{}

func (mr *MapReducer) Map(in chan interface{}, out chan interface{}) {
	for v := range in {
		out <- v
	}
}

func (mr *MapReducer) Reduce(in chan interface{}) interface{} {
	res := 0
	for v := range in {
		res += v.(int)
	}
	return res
}

func TestMapReduce(t *testing.T) {

	conf := NewMapReduceConfig()

	// Feed input channel
	go func(in chan interface{}) {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}(conf.InChan)

	mr := &MapReducer{}

	result, _ := Run(mr, conf)

	if result != 4950 {
		t.Fail()
	}
}
