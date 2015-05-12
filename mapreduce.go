package mapreduce

type Map interface {
	Output(chan interface{})
	Input(chan interface{})
	Process()
}

type Reduce interface {
	Input(chan interface{})
	Result() interface{}
}

type Configuration struct {
	mapperCount int
}

func newMapReduce(m Map, r Reduce, config Configuration) error {

	return nil
}
