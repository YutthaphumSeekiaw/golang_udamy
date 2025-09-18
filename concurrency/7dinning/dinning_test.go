package main

import (
	"testing"
	"time"
)

func TestDinning(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished := []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("Expected 5 philosophers to finish, but got %d", len(orderFinished))
		}
	}
}

func Test_dine(t *testing.T) {
	var thetest = []struct {
		name  string
		delay time.Duration
	}{
		{name: "no delay", delay: 0 * time.Millisecond},
		{name: "with delay", delay: 1 * time.Millisecond},
		{name: "long delay", delay: 3 * time.Millisecond},
	}

	for _, tt := range thetest {
		orderFinished := []string{}
		eatTime = tt.delay
		thinkTime = tt.delay
		sleepTime = tt.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s Expected 5 philosophers to finish, but got %d", tt.name, len(orderFinished))
		}
	}

}
