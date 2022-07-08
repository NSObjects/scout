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

type WorkerPool interface {
	Run()
	AddTask(task func())
	Stop()
}

type work struct {
	workerCount int
	taskChan    chan func()
	close       chan struct{}
}

func NewWorkerPool(workerCount int) WorkerPool {
	return &work{
		workerCount: workerCount,
		taskChan:    make(chan func()),
		close:       make(chan struct{}),
	}
}

func (w *work) Run() {
	for i := 0; i < w.workerCount; i++ {
		go func() {
			for {
				select {
				case <-w.close:
					return
				case task := <-w.taskChan:
					task()
				}
			}
		}()
	}
}

func (w *work) AddTask(task func()) {
	w.taskChan <- task
}

func (w work) Stop() {
	w.close <- struct{}{}
}
