package maps

import (
	"github.com/goextension/mapper/contacts"
	"github.com/goextension/mapper/sort"
)

// MakeMapper Get Mappable instance.
func MakeMapper[K string | int, V any]() contacts.Mappable[K, V] {
	return &Mapper[K, V]{
		maps:     make(map[K]V),
		sortable: &sort.MapSorter[K, V]{},
	}
}
