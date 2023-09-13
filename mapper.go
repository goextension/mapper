package mapper

import "github.com/goextension/mapper/contacts"

type Mapper[K string | int, V any] struct {
	maps map[K]V
}

func (mapper *Mapper[K, V]) SetMap(haystack K, needle V) contacts.Mappable[K, V] {

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

	carry := make([]K, 0, len(mapper.maps))

	mapper.Each(func(value V, key K) {
		carry = append(carry, key)
	})

	return carry
}

func (mapper *Mapper[K, V]) Values() []V {

	values := make([]V, 0, len(mapper.maps))

	mapper.Each(func(value V, key K) {
		values = append(values, value)
	})

	return values
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

func (mapper *Mapper[K, V]) Map(closure func(value V, key K, mappable contacts.Mappable[K, V])) contacts.Mappable[K, V] {

	mapper.Each(func(value V, key K) {
		closure(value, key, mapper)
	})

	return mapper
}

func (mapper *Mapper[K, V]) Each(closure func(value V, key K)) {
	for key, value := range mapper.maps {
		closure(value, key)
	}
}

func (mapper *Mapper[K, V]) Flush() {
	mapper.maps = make(map[K]V)
}
