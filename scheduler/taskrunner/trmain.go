package taskrunner

import "time"

//timer
//setup
//start{trigger->task->runner}

//timer,task,runner,longLived

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select { //unblock
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// Start video file cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
