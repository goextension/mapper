package sort

import "github.com/goextension/mapper/contacts/maps"

type Sortable[K maps.Argument, V any] interface {

	// SortBy Performs a sort operation on a given map.
	SortBy() maps.Mappable[K, V]

	SortByDesc() maps.Mappable[K, V]
}
