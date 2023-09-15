package sort

import (
	"github.com/goextension/mapper/contacts"
	"slices"
	"strings"
)

type MapSorter[K string | int, V any] struct {
	caches []K

	sorted bool

	mappable contacts.Mappable[K, V]

	trigger string
}

func (sorter *MapSorter[K, V]) Each(closure func(key K)) {
	for _, key := range sorter.caches {
		closure(key)
	}
}

func (sorter *MapSorter[K, V]) SortBy() contacts.Mappable[K, V] {

	sorter.fireEvent("asc")

	return sorter.sortSync()
}

func (sorter *MapSorter[K, V]) SortByDesc() contacts.Mappable[K, V] {

	sorter.fireEvent("desc")

	return sorter.sortSync()
}

func (sorter *MapSorter[K, V]) fireEvent(trigger string) {
	sorter.trigger = trigger
}

func (sorter *MapSorter[K, V]) sortSync() contacts.Mappable[K, V] {

	if !sorter.IsSorted() {

		sorter.caches = make([]K, 0, sorter.mappable.Count())

		sorter.store(sorter.getSortValues())

		sorter.sorted = true

		return sorter.mappable
	}

	if !slices.Equal(sorter.caches, sorter.mappable.Keys()) {

		// 如果缓存与mappable不同，则重新赋值去重
	}

	sorter.store(sorter.getSortValues())

	return sorter.mappable

}

func (sorter *MapSorter[K, V]) getSortValues() []K {

	sortValues := sorter.mappable.Keys()

	if sorter.isDesc() {
		slices.Reverse(sortValues)

		return sortValues
	}

	slices.Sort(sortValues)

	return sortValues
}

func (sorter *MapSorter[K, V]) isDesc() bool {
	return strings.EqualFold(sorter.trigger, "desc")
}

func (sorter *MapSorter[K, V]) store(values []K) {
	for _, key := range values {
		sorter.caches = append(sorter.caches, key)
	}
}

func (sorter *MapSorter[K, V]) SetMappable(mappable contacts.Mappable[K, V]) contacts.Sortable[K, V] {

	if sorter.mappable == nil {
		sorter.mappable = mappable
	}

	return sorter
}

func (sorter *MapSorter[K, V]) IsSorted() bool {
	return sorter.sorted
}
