package sort

import "github.com/goextension/mapper/contacts/maps"

type Sorter[K maps.Argument, V any] interface {
	Sortable[K, V]

	SetMappable(mappable maps.Mappable[K, V]) Sortable[K, V]

	GetSortValues() []K
}
