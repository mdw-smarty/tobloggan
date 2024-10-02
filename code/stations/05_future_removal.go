package stations

import (
	"time"

	"tobloggan/code/contracts"
)

type FutureRemoval struct {
	now time.Time
}

func NewFutureRemoval(now time.Time) contracts.Station {
	return &FutureRemoval{now: now}
}

func (this *FutureRemoval) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		if !input.Date.After(this.now) {
			output(input)
		}
	default:
		output(input)
	}
}
