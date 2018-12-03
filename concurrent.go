package gutil

import "reflect"

type ConcurrentTask struct {
	Value interface{}
	Do    func() interface{}
}

func Concurrent(tasks []*ConcurrentTask) int {
	cases := make([]reflect.SelectCase, len(tasks))
	for i, task := range tasks {
		c := make(chan interface{})
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)}
		go func(t *ConcurrentTask) {
			r := t.Do()
			c <- r
		}(task)
	}

	counter := 0
	completeCounter := 0
	for counter < len(tasks) {
		chosen, value, ok := reflect.Select(cases)
		counter++
		if ok {
			completeCounter++
		}
		tasks[chosen].Value = value.Interface()
	}
	return completeCounter
}
