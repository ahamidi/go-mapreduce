package mapreduce

type MapReduce interface {
	Map(in chan interface{}, out chan interface{})
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

	// Map
	go mr.Map(mr.Config.inChan, mr.Config.outChan)

	// Reduce
	result := mr.Reduce(mr.Config.outChan)

	return result, nil
}

func newMapReducer(config Configuration) *MapReducer {

	return &MapReducer{
		Config: config,
	}
}
