package timeproblem

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeToAngle(t *testing.T) {
	tests := []struct {
		hours   int
		minutes int
		result  int
		err     error
	}{
		{hours: 0, minutes: 0, result: 0, err: nil},
		{hours: 0, minutes: 1, result: 6, err: nil},
		{hours: 0, minutes: 59, result: 6, err: nil},
		{hours: 0, minutes: 15, result: 90, err: nil},
		{hours: 0, minutes: 45, result: 90, err: nil},
		{hours: 0, minutes: 5, result: 30, err: nil},
		{hours: 0, minutes: 55, result: 30, err: nil},
		{hours: 1, minutes: 0, result: 30, err: nil},
		{hours: 1, minutes: 1, result: 24, err: nil},
		{hours: 1, minutes: 59, result: 36, err: nil},
		{hours: 1, minutes: 15, result: 60, err: nil},
		{hours: 1, minutes: 45, result: 120, err: nil},
		{hours: 1, minutes: 5, result: 0, err: nil},
		{hours: 1, minutes: 55, result: 60, err: nil},
		{hours: 6, minutes: 0, result: 180, err: nil},
		{hours: 6, minutes: 1, result: 174, err: nil},
		{hours: 6, minutes: 59, result: 174, err: nil},
		{hours: 11, minutes: 0, result: 30, err: nil},
		{hours: -1, minutes: 0, result: 0, err: errors.New("Unexpected hours = -1")},
		{hours: 12, minutes: 0, result: 0, err: errors.New("Unexpected hours = 12")},
		{hours: 0, minutes: -1, result: 0, err: errors.New("Unexpected minutes = -1")},
		{hours: 0, minutes: 60, result: 0, err: errors.New("Unexpected minutes = 60")},
	}

	for _, tt := range tests {
		result, err := TimeToAngle(tt.hours, tt.minutes)

		assert.Equal(t, tt.result, result, fmt.Sprintf("hours=%d, minures=%d", tt.hours, tt.minutes))
		assert.Equal(t, tt.err, err, fmt.Sprintf("hours=%d, minures=%d", tt.hours, tt.minutes))
	}
}
