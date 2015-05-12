package mapreduce

import (
	"log"
	"testing"
)

type testMap struct {
	inChan  chan interface{}
	outChan chan interface{}
}

func (tm *testMap) Input(ch chan interface{}) {
	tm.inChan = ch
}

func (tm *testMap) Output(ch chan interface{}) {
	tm.outChan = ch
}

func (tm *testMap) Process() {
	go func() {
		for v := range tm.inChan {
			tm.outChan <- v
		}
		close(tm.outChan)
	}()
}

func TestMap(t *testing.T) {
	inChan := make(chan interface{})
	outChan := make(chan interface{})

	tm := &testMap{}
	tm.Input(inChan)
	tm.Output(outChan)

	go func(ch chan interface{}) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(inChan)

	tm.Process()

	for v := range outChan {
		log.Println(v)
	}

	t.Fail()
}

func TestReduce(t *testing.T) {

}

func TestMapReduce(t *testing.T) {

}
