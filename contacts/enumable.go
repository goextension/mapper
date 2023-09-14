package contacts

type Enumerable[K string | int, V any] interface {
	Filter(closure func(value V, key K) bool) Mappable[K, V]

	Map(closure func(value V, key K, maps map[K]V)) Mappable[K, V]

	Each(closure func(value V, key K))
}
