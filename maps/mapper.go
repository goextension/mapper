package maps

import (
	"github.com/goextension/mapper/contacts/maps"
	"github.com/goextension/mapper/contacts/sort"
	"sync"
)

type Mapper[K maps.Argument, V any] struct {
	maps map[K]V

	sortable sort.Sorter[K, V]

	mutex sync.RWMutex
}

func (mapper *Mapper[K, V]) Store(haystack K, needle V) maps.Mappable[K, V] {

	mapper.mutex.Lock()

	mapper.maps[haystack] = needle

	mapper.mutex.Unlock()

	return mapper
}

func (mapper *Mapper[K, V]) Equal(haystack K, closure func(value V) bool) bool {

	value, exists := mapper.maps[haystack]

	if !exists {
		return false
	}

	return closure(value)
}

func (mapper *Mapper[K, V]) DeleteWithCopy(haystack K) V {

	value := mapper.Get(haystack)

	mapper.Unset(haystack)

	return value
}

func (mapper *Mapper[K, V]) Get(haystack K) V {

	mapper.mutex.RLock()

	defer mapper.mutex.RUnlock()

	return mapper.maps[haystack]
}

func (mapper *Mapper[K, V]) Unset(haystack K) {

	if mapper.Has(haystack) {
		delete(mapper.maps, haystack)
	}
}

func (mapper *Mapper[K, V]) Has(haystack K) bool {

	_, exists := mapper.maps[haystack]

	return exists
}

func (mapper *Mapper[K, V]) Keys() []K {
	return mapper.sortable.GetSortValues()
}

func (mapper *Mapper[K, V]) Values() []V {

	carry := make([]V, 0, mapper.Count())

	for _, key := range mapper.sortable.GetSortValues() {
		carry = append(carry, mapper.Get(key))
	}

	return carry
}

func (mapper *Mapper[K, V]) IsEmpty() bool {
	return !mapper.IsNotEmpty()
}

func (mapper *Mapper[K, V]) IsNotEmpty() bool {
	return mapper.Count() > 0
}

func (mapper *Mapper[K, V]) Filter(closure func(value V, key K) bool) maps.Mappable[K, V] {

	mapper.Each(func(value V, key K) {
		if closure(value, key) {
			mapper.Unset(key)
		}
	})

	return mapper
}

func (mapper *Mapper[K, V]) Map(closure func(value V, key K, maps map[K]V)) maps.Mappable[K, V] {

	mapper.Each(func(value V, key K) {
		closure(value, key, mapper.maps)
	})

	return mapper
}

func (mapper *Mapper[K, V]) Each(closure func(value V, key K)) {

	mapper.mutex.RLock()

	for key, value := range mapper.maps {
		closure(value, key)
	}

	mapper.mutex.RUnlock()
}

func (mapper *Mapper[K, V]) SortBy() maps.Mappable[K, V] {
	return mapper.sortable.SetMappable(mapper).SortBy()
}

func (mapper *Mapper[K, V]) SortByDesc() maps.Mappable[K, V] {
	return mapper.sortable.SetMappable(mapper).SortByDesc()
}

func (mapper *Mapper[K, V]) Count() int {
	return len(mapper.maps)
}

func (mapper *Mapper[K, V]) Flush() {
	mapper.maps = make(map[K]V)
}
