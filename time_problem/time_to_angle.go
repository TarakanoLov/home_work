package timeproblem

import (
	"fmt"
)

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func TimeToAngle(hours int, minutes int) (int, error) {
	if hours < 0 || hours > 11 {
		return 0, fmt.Errorf("Unexpected hours = %d", hours)
	}
	if minutes < 0 || minutes > 59 {
		return 0, fmt.Errorf("Unexpected minutes = %d", minutes)
	}
	angle := abs(hours*5-minutes) * 6
	return min(angle, 360-angle), nil
}
