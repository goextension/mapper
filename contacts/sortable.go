package contacts

type Sortable[K string | int, V any] interface {

	// SortBy Performs a sort operation on a given map.
	SortBy() Mappable[K, V]

	SortByDesc() Mappable[K, V]
}
