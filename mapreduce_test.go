package mapreduce

import (
	"sync"
	"testing"
)

type MapReducer struct{}

func (mr *MapReducer) Map(in chan interface{}, out chan interface{}, wg *sync.WaitGroup) {
	for v := range in {
		out <- v
	}
	wg.Done()
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

	conf := NewMapReduceConfig()
	conf.InChan = inChan
	conf.OutChan = outChan
	conf.MapperCount = 3

	// Feed input channel
	go func(in chan interface{}) {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}(inChan)

	mr := &MapReducer{}

	result, _ := Run(mr, conf)

	if result != 4950 {
		t.Fail()
	}
}
