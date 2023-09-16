package observer

import "strings"

type SortEventObserver struct {
	event string

	sorted bool
}

func (observer *SortEventObserver) FireEvent(event string) {
	observer.event = event

	if !observer.sorted {
		observer.sorted = true
	}
}

func (observer *SortEventObserver) HasEvent(event string) bool {
	return strings.EqualFold(observer.event, event)
}

func (observer *SortEventObserver) IsTrigger() bool {
	return observer.sorted
}
