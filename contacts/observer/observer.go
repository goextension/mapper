package observer

type Observer interface {
	FireEvent(event string)

	Is(event string) bool

	HasEvent() bool
}
