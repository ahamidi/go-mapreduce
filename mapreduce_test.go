package mapreduce

import (
	"log"
	"testing"
)

type testMapReduceFunctions struct{}

func (tmr *testMapReduceFunctions) Map(in chan interface{}, out chan interface{}) {
	values := []int{}
	for v := range in {
		values = append(values, v.(int))
	}
	log.Println("Values:", values)

	for v := range values {
		out <- v
	}
	close(out)
}

func (tmr *testMapReduceFunctions) Reduce(in chan interface{}) interface{} {
	res := 0
	for v := range in {
		res += v.(int)
	}
	return res
}

func TestMap(t *testing.T) {
	inChan := make(chan interface{})
	outChan := make(chan interface{})

	tm := &testMapReduceFunctions{}

	config := Configuration{
		mapperCount: 1,
		inChan:      inChan,
		outChan:     outChan,
	}

	// Feed input channel
	go func(in chan interface{}) {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}(inChan)

	result, _ := MapReduce(tm, config)
	log.Println("Result:", result)

	t.Fail()
}

func TestReduce(t *testing.T) {

}

func TestMapReduce(t *testing.T) {

}
