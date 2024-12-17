package worker_pool

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(10, 100)

	assert.NotNil(t, wp.taskQueue)
	assert.NotNil(t, wp.stopCh)
	assert.False(t, wp.stopAllTasks)
}

func TestSubmit(t *testing.T) {
	wp := NewWorkerPool(10, 10)

	startTime := time.Now()
	var timeOfComplete time.Time
	assert.True(t, wp.Submit(func() {
		time.Sleep(time.Second)
		timeOfComplete = time.Now()
	}))
	assert.Less(t, time.Since(startTime), time.Second)
	assert.Equal(t, time.Time{}, timeOfComplete)

	wp.StopWait()
	assert.LessOrEqual(t, startTime.Add(time.Second), timeOfComplete)
}

func TestSubmitWait(t *testing.T) {
	wp := NewWorkerPool(10, 10)

	startTime := time.Now()
	var timeOfComplete time.Time
	assert.True(t, wp.SubmitWait(func() {
		time.Sleep(time.Second)
		timeOfComplete = time.Now()
	}))
	assert.LessOrEqual(t, time.Second, time.Since(startTime))
	assert.LessOrEqual(t, startTime.Add(time.Second), timeOfComplete)
}

func TestStop(t *testing.T) {
	wp := NewWorkerPool(10, 10000)

	var ops atomic.Uint64

	for i := 0; i < 50; i++ {
		assert.True(t, wp.Submit(func() {
			time.Sleep(time.Second)
			ops.Add(1)
		}))
	}

	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, ops.Load(), uint64(0))
	wp.Stop()

	assert.Equal(t, ops.Load(), uint64(10))
}

func TestStopWait(t *testing.T) {
	wp := NewWorkerPool(10, 10000)

	var ops atomic.Uint64

	for i := 0; i < 50; i++ {
		assert.True(t, wp.Submit(func() {
			time.Sleep(time.Second)
			ops.Add(1)
		}))
	}

	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, ops.Load(), uint64(0))
	wp.StopWait()

	assert.Equal(t, ops.Load(), uint64(50))
}
