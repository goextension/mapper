package sort

import (
	"github.com/goextension/mapper/contacts/maps"
	"github.com/goextension/mapper/contacts/observer"
	"github.com/goextension/mapper/contacts/sort"
	"slices"
)

type MapSorter[K string | int, V any] struct {
	observer observer.Observer
	mappable maps.Mappable[K, V]
}

func (sorter *MapSorter[K, V]) GetSortValues() []K {

	if !sorter.observer.HasEvent() {
		return sorter.mappable.Keys()
	}

	values := sorter.mappable.Keys()

	if sorter.observer.Is("desc") {
		slices.Reverse(values)
		return values
	}

	slices.Sort(values)
	return values
}

func (sorter *MapSorter[K, V]) SortBy() maps.Mappable[K, V] {

	sorter.observer.FireEvent("asc")

	return sorter.mappable
}

func (sorter *MapSorter[K, V]) SortByDesc() maps.Mappable[K, V] {

	sorter.observer.FireEvent("desc")

	return sorter.mappable
}

func (sorter *MapSorter[K, V]) SetMappable(mappable maps.Mappable[K, V]) sort.Sortable[K, V] {

	if sorter.mappable == nil {
		sorter.mappable = mappable
	}

	return sorter
}
