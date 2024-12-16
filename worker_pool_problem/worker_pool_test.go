package worker_pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerPoolNew(t *testing.T) {
	wp := NewWorkerPool(10, 100)

	assert.NotNil(t, wp.taskQueue)
	assert.NotNil(t, wp.wg)
	assert.NotNil(t, wp.stopCh)
	assert.False(t, wp.stopAllTasks)
}
