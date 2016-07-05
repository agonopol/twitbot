package lib

import (
	"math/rand"
	"time"
)

type Task interface {
	Execute()
}

func Schedule(task Task, min, max int64) func() {
	return func() {
		for {
			duration := time.Duration(rand.Int63n(max-min) + min)
			time.Sleep(duration)
			task.Execute()
		}
	}
}
