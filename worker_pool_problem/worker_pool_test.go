package worker_pool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(10, 100)

	assert.NotNil(t, wp.taskQueue)
	assert.NotNil(t, wp.wg)
	assert.NotNil(t, wp.stopCh)
	assert.False(t, wp.stopAllTasks)
}

func TestSubmit(t *testing.T) {
	wp := NewWorkerPool(10, 10)

	startTime := time.Now()
	var timeOfComplete time
	assert.True(t, wp.Submit(func() {
		time.Sleep(time.Second)
		timeOfComplete = time.Now()
	}))
	assert.Less(t, time.Since(startTime), time.Second)
	assert.LessOrEqual(t, time.Second, timeOfComplete-startTime)
}
