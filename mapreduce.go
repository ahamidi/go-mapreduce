package mapreduce

import (
	"sync"
)

type MapReduce interface {
	Map(in chan interface{}, out chan interface{}, wg *sync.WaitGroup)
	Reduce(in chan interface{}) interface{}
}

type MapReducer struct {
	Config Configuration
}

type Configuration struct {
	mapperCount int
	inChan      chan interface{}
	outChan     chan interface{}
}

func (mr *MapReducer) Run() (interface{}, error) {

	var wg sync.WaitGroup

	// Map
	for i := 0; i < mr.Config.mapperCount; i++ {
		wg.Add(1)
		go mr.Map(mr.Config.inChan, mr.Config.outChan, &wg)
	}

	go func(w *sync.WaitGroup) {
		w.Wait()
		close(mr.Config.outChan)
	}(&wg)

	// Reduce
	resultChan := make(chan interface{}, 1)
	go func(res chan interface{}) {
		res <- mr.Reduce(mr.Config.outChan)
	}(resultChan)

	return <-resultChan, nil
}

func newMapReducer(config Configuration) *MapReducer {

	return &MapReducer{
		Config: config,
	}
}
