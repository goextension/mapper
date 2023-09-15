package contacts

type Sorter[K string | int, V any] interface {
	Sortable[K, V]

	SetMappable(mappable Mappable[K, V]) Sortable[K, V]

	Each(closure func(key K))

	IsSorted() bool
}
