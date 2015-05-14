package mapreduce

import (
	"fmt"
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

type ExampleMapReducer struct{}

func (mr *ExampleMapReducer) Map(in chan interface{}, out chan interface{}) {
	for v := range in {
		out <- v
	}
}

// Example Map Function
func (mr *ExampleMapReducer) Example_map(in chan interface{}, out chan interface{}) {
	for v := range in {
		out <- v
	}
}

func (mr *ExampleMapReducer) Reduce(in chan interface{}) interface{} {
	res := 0
	for v := range in {
		res += v.(int)
	}
	return res
}

// Example Reduce Function
func (mr *ExampleMapReducer) Example_reduce(in chan interface{}) interface{} {
	res := 0
	for v := range in {
		res += v.(int)
	}
	return res
}

func Example() {
	conf := NewMapReduceConfig()

	// Feed input channel
	go func(in chan interface{}) {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}(conf.InChan)

	var mr ExampleMapReducer

	result, _ := Run(&mr, conf)

	fmt.Println(result)
	// Output:
	// 4950

}
