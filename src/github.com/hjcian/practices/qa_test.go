package main

import (
	"testing"
)

func TestCalcScore(t *testing.T) {
	userAns := []int{1, 2, 3, 45, 6}
	examAns := []int{1, 2, 3, 45, 6}
	score := calcScore(&userAns, &examAns)
	if score != 5 {
		t.Errorf("should be 5 by given: %v and %v", userAns, examAns)
	}
	userAns = []int{1, 2}
	examAns = []int{1, 2}
	score = calcScore(&userAns, &examAns)
	if score != 2 {
		t.Errorf("should be 1 by given: %v and %v", userAns, examAns)
	}
}
