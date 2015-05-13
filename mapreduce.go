package mapreduce

import (
	"sync"
)

type MapReduce interface {
	Map(in chan interface{}, out chan interface{})
	Reduce(in chan interface{}) interface{}
}

type Configuration struct {
	MapperCount int
	InChan      chan interface{}
	OutChan     chan interface{}
}

func NewMapReduceConfig() *Configuration {
	return &Configuration{}
}

func Run(mr MapReduce, c *Configuration) (interface{}, error) {

	var wg sync.WaitGroup

	// Map
	for i := 0; i < c.MapperCount; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			mr.Map(c.InChan, c.OutChan)
			wg.Done()
		}(&wg)
	}

	go func(w *sync.WaitGroup) {
		w.Wait()
		close(c.OutChan)
	}(&wg)

	// Reduce
	resultChan := make(chan interface{}, 1)
	go func(res chan interface{}) {
		res <- mr.Reduce(c.OutChan)
	}(resultChan)

	return <-resultChan, nil
}
