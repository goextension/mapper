package contacts

type Enumerable[K string | int, V any] interface {
	Filter(closure func(value V, key K) bool) Mappable[K, V]

	Map(closure func(value V, key K, mappable Mappable[K, V])) Mappable[K, V]

	Each(closure func(value V, key K))
}
