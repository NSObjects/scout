/*
 *
 * work.go
 * workpool
 *
 * Created by lintao on 2022/7/8 10:38 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package workpool

import "sync"

type WorkerPool interface {
	AddTask(task func())
	Stop()
	Wait()
}

type work struct {
	workerCount int
	taskChan    chan func()
	close       chan struct{}
	wg          sync.WaitGroup
}

func NewWorkerPool(workerCount int) WorkerPool {
	wk := &work{
		workerCount: workerCount,
		taskChan:    make(chan func()),
		close:       make(chan struct{}),
	}
	wk.run()
	return wk
}

func (w *work) run() {
	for i := 0; i < w.workerCount; i++ {
		go func() {
			for {
				select {
				case <-w.close:
					return
				case task := <-w.taskChan:
					task()
					w.wg.Done()
				}
			}
		}()
	}
}

func (w *work) Wait() {
	w.wg.Wait()
	w.close <- struct{}{}
}

func (w *work) AddTask(task func()) {
	w.wg.Add(1)
	w.taskChan <- task
}

func (w work) Stop() {
	w.close <- struct{}{}
}
