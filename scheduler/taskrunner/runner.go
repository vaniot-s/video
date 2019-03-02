package taskrunner

//生产消费者 常驻
//runner
//startDispatcher
//control channel dispatcher executor交换信息
//data channel  数据
type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn //func
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

//常驻任务
func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived { //是否常驻
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select { //no block
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data) //写数据 写完或 写入错误 返回err
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE //
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data) //取数据
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH //预置
	r.startDispatch()
}
