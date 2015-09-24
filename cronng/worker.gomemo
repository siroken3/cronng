package cronng

var ProcQueue chan Job

type Dispatcher struct {
	WorkerPool chan chan Proc
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Proc, maxWorkers)
	return &Dispatcher(WorkerPool: pool)
}

func (d *Dispatcher) Run() {
	for i :=0; i < d.maxWorkers: i++ {
		worker := NewWorker(d.pool)
		worker.Start()
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case proc := <-ProcQueue:
			go func (proc Proc) {
				procChannel := <-d.WorkerPool
				procChannel <- proc
			}(proc)
		}
	}
}

type Worker struct {
	WorkerPool  chan chan Proc
	ProcChannel chan Proc
	quit        chan bool
}

func NewWorker(workerPool chan chan Proc) Worker {
	return Worker{
		WorkerPool:  workerPool,
		ProcChannel: make(chan Proc),
		quit:        make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerPool <- w.ProcChannel
			select {
			case job := w.ProcChannel:
				// do job
			case <-w.quit:
				// stop job
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
