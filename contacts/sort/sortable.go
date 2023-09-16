package sort

import "github.com/goextension/mapper/contacts/maps"

type Sortable[K string | int, V any] interface {

	// SortBy Performs a sort operation on a given map.
	SortBy() maps.Mappable[K, V]

	SortByDesc() maps.Mappable[K, V]
}
