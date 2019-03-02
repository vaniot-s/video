package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	VIDEO_PATH = "./videos/"
)

type controlChan chan string //control

type dataChan chan interface{} //go 泛型的机制

type fn func(dc dataChan) error //func
