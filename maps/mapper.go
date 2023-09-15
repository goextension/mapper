package maps

import "github.com/goextension/mapper/contacts"

type Mapper[K string | int, V any] struct {
	maps map[K]V

	sortable contacts.Sorter[K, V]
}

func (mapper *Mapper[K, V]) Store(haystack K, needle V) contacts.Mappable[K, V] {

	mapper.maps[haystack] = needle

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

	carry := make([]K, 0, mapper.Count())

	if mapper.sortable.IsSorted() {
		mapper.sortable.Each(func(key K) {
			carry = append(carry, key)
		})

		return carry
	}

	mapper.Each(func(value V, key K) {
		carry = append(carry, key)
	})

	return carry
}

func (mapper *Mapper[K, V]) Values() []V {

	carry := make([]V, 0, mapper.Count())

	if mapper.sortable.IsSorted() {

		mapper.sortable.Each(func(key K) {
			carry = append(carry, mapper.Get(key))
		})

		return carry
	}

	mapper.Each(func(value V, key K) {
		carry = append(carry, value)
	})

	return carry
}

func (mapper *Mapper[K, V]) IsEmpty() bool {
	return !mapper.IsNotEmpty()
}

func (mapper *Mapper[K, V]) IsNotEmpty() bool {
	return len(mapper.maps) > 0
}

func (mapper *Mapper[K, V]) Filter(closure func(value V, key K) bool) contacts.Mappable[K, V] {

	mapper.Each(func(value V, key K) {
		if closure(value, key) {
			mapper.Unset(key)
		}
	})

	return mapper
}

func (mapper *Mapper[K, V]) Map(closure func(value V, key K, maps map[K]V)) contacts.Mappable[K, V] {

	mapper.Each(func(value V, key K) {
		closure(value, key, mapper.maps)
	})

	return mapper
}

func (mapper *Mapper[K, V]) Each(closure func(value V, key K)) {
	for key, value := range mapper.maps {
		closure(value, key)
	}
}

func (mapper *Mapper[K, V]) SortBy() contacts.Mappable[K, V] {
	return mapper.sortable.SetMappable(mapper).SortBy()
}

func (mapper *Mapper[K, V]) SortByDesc() contacts.Mappable[K, V] {
	return mapper.sortable.SetMappable(mapper).SortByDesc()
}

func (mapper *Mapper[K, V]) Count() int {
	return len(mapper.maps)
}

func (mapper *Mapper[K, V]) Flush() {
	mapper.maps = make(map[K]V)
}
