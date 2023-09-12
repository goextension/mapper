package mapper

import "goextension/mapper/contacts"

type Mapper[K string | int, V any] struct {
	maps map[K]V
}

func (mapper *Mapper[K, V]) SetMap(haystack K, needle V) contacts.Mappable[K, V] {

	mapper.maps[haystack] = needle

	return mapper
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

	keys := make([]K, len(mapper.maps))

	for key, _ := range mapper.maps {
		keys = append(keys, key)
	}

	return keys
}

func (mapper *Mapper[K, V]) Values() []V {

	values := make([]V, len(mapper.maps))

	for _, value := range mapper.maps {
		values = append(values, value)
	}

	return values
}

func (mapper *Mapper[K, V]) IsEmpty() bool {
	return !mapper.IsNotEmpty()
}

func (mapper *Mapper[K, V]) IsNotEmpty() bool {
	return len(mapper.maps) > 0
}

func (mapper *Mapper[K, V]) Flush() {
	mapper.maps = make(map[K]V)
}
