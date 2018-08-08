package pool


// 定义任务
type Job func()

// 定义工人，具体执行任务的单元
type worker struct {
	// 任务，源源不断的任务
	job chan Job
	// 属于哪个工作池的工人
	workerPool *WorkerPool
	// 下班标志
	quit chan bool
}

// 开始干活了
func (w worker) start() {
	go func() {
		// 工人不断的做任务。。。
		for {
			// 没事干，到工作池里面等待任务
			w.workerPool.worker <- &w
			select {
			// 任务来了
			case job := <-w.job:
				job()
			// 下班了
			case <-w.quit:
				// 走人
				return
			}
		}
	}()
}

// 准备下班了
func (w worker) stop() {
	w.quit <- true
}

// 招聘一个工人
func newWorker(workerPool *WorkerPool) *worker {
	return &worker{
		job:        make(chan Job),
		workerPool: workerPool,
		quit:       make(chan bool),
	}
}

// 工人池
type WorkerPool struct {
	// 任务队列
	jobQueue chan Job
	// 工人池里面闲置的工人
	worker chan *worker
	// 任务队列容量
	jobSize int
	// 工人数目
	workerNum int
	// 池中的工人数
	workerNumCount int
	// 解散标志
	quit chan bool
}

// 启动工作池，派遣任务给工人
func (wp WorkerPool) Start() {
	go func() {
		for {
			select {
			// 从任务队列里取出任务
			case job := <-wp.jobQueue:
				// 从工人池里面找到一个很闲的工人
				worker := <-wp.worker
				// 把任务派遣给他
				worker.job <- job
			// 下班了
			case <-wp.quit:
				for i := 0; i < wp.workerNum; i++ {
					// 从工人池里面找到一个很闲的工人
					worker := <-wp.worker
					// 告知工人下班了
					worker.stop()
				}
				return
			}
		}
	}()
}

// 来任务了，加入到任务队列，会等待任务队列空闲
func (wp WorkerPool) Add(job Job) {
	// 任务放入任务队列
	wp.jobQueue <- job
}

// 工人池解散了
func (wp WorkerPool) Stop() {
	wp.quit <- true
}

// 创建一个工作池
func NewWorkerPool(workerNum int, jobSize int) *WorkerPool {
	wp := &WorkerPool{
		jobQueue:  make(chan Job, jobSize),
		worker:    make(chan *worker, workerNum),
		jobSize:   jobSize,
		workerNum: workerNum,
		quit:      make(chan bool),
	}
	// 初始化工人
	for i := 0; i < workerNum; i++ {
		worker := newWorker(wp)
		worker.start()
	}
	return wp
}
