package observer

type Observer interface {
	FireEvent(event string)

	HasEvent(event string) bool

	IsTrigger() bool
}
