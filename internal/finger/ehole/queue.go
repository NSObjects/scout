/*
 *
 * queue.go
 * finger
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package ehole

import (
	"container/list"
	"sync"
)

type Queue struct {
	data *list.List
	lock sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		data: new(list.List),
	}
}

func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	var value interface{}
	if element := q.data.Back(); element != nil {
		value = element.Value
		q.data.Remove(element)
	}

	return value
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.data.Len()
}

func (q *Queue) Push(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data.PushFront(v)
}
