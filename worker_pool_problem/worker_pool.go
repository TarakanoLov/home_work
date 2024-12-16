package worker_pool

import (
	"sync"
)

type WorkerPool struct {
	taskQueue    chan func()
	wg           sync.WaitGroup
	stopCh       chan struct{}
	m            sync.Mutex
	stopAllTasks bool
}

func startWorker(wp *WorkerPool) {
	for {
		select {
		case <-wp.stopCh:
			return
		default:
			task := <-wp.taskQueue
			if task != nil {
				task()
				wp.wg.Done()
			} else {
				return
			}
		}
	}
}

func NewWorkerPool(numberOfWorkers int, queueSize int) *WorkerPool {
	wp := &WorkerPool{
		taskQueue:    make(chan func(), queueSize),
		stopCh:       make(chan struct{}),
		stopAllTasks: false,
	}

	for i := 0; i < numberOfWorkers; i++ {
		go startWorker(wp)
	}
	return wp
}

// Submit - добавить таску в воркер пул
func (wp *WorkerPool) Submit(task func()) bool {
	wp.m.Lock()
	if wp.stopAllTasks {
		wp.m.Unlock()
		return false
	}
	wp.wg.Add(1)
	wp.taskQueue <- task
	wp.m.Unlock()
	return true
}

// SubmitWait - добавить таску в воркер пул и дождаться окончания ее выполнения
func (wp *WorkerPool) SubmitWait(task func()) bool {
	wp.m.Lock()
	if wp.stopAllTasks {
		wp.m.Unlock()
		return false
	}
	ch := make(chan struct{})
	wp.wg.Add(1)

	wp.taskQueue <- func() {
		task()
		ch <- struct{}{}
	}

	wp.m.Unlock()

	<-ch
	return true
}

// Stop - остановить воркер пул, дождаться выполнения только тех тасок,
// которые выполняются сейчас
func (wp *WorkerPool) Stop() {
	wp.m.Lock()
	wp.stopAllTasks = true
	wp.m.Unlock()

	close(wp.stopCh)
	close(wp.taskQueue)
	wp.wg.Wait()
}

// StopWait - остановить воркер пул, дождаться выполнения всех тасок,
// даже тех, что не начали выполняться, но лежат в очереди
func (wp *WorkerPool) StopWait() {
	wp.m.Lock()
	wp.stopAllTasks = true
	wp.m.Unlock()

	wp.wg.Wait()
}
