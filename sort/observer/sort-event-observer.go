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

func (observer *SortEventObserver) Is(event string) bool {
	return strings.EqualFold(observer.event, event)
}

func (observer *SortEventObserver) HasEvent() bool {
	return observer.sorted
}
