package mapreduce

type MapReduceFunctions interface {
	Map(in chan interface{}, out chan interface{})
	Reduce(in chan interface{}) interface{}
}

type Configuration struct {
	mapperCount int
	inChan      chan interface{}
	outChan     chan interface{}
}

func MapReduce(mr MapReduceFunctions, config Configuration) (interface{}, error) {

	// Map
	go mr.Map(config.inChan, config.outChan)

	// Reduce
	result := mr.Reduce(config.outChan)

	return result, nil
}
