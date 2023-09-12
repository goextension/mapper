package mapper

import `goextension/contacts`

type Mapper[K string | int] struct {
	maps map[K]any
}

func (mapper *Mapper[K]) SetMap(haystack K, needle any) contacts.Mappable[K] {

	if mapper.maps == nil {
		mapper.maps = make(map[K]any)
	}

	mapper.maps[haystack] = needle

	return mapper
}

func (mapper *Mapper[K]) Get(haystack K) any {

	if mapper.HasKey(haystack) {
		return mapper.maps[haystack]
	}

	return nil
}

func (mapper *Mapper[K]) Delete(haystack K) {

	if mapper.HasKey(haystack) {
		delete(mapper.maps, haystack)
	}

}

func (mapper *Mapper[K]) HasKey(haystack K) bool {

	_, exists := mapper.maps[haystack]

	return exists
}

func (mapper *Mapper[K]) Keys() []K {

	keys := make([]K, len(mapper.maps))

	for key, _ := range mapper.maps {
		keys = append(keys, key)
	}

	return keys
}

func (mapper *Mapper[K]) Values() []any {

	values := make([]any, len(mapper.maps))

	for _, value := range mapper.maps {
		values = append(values, value)
	}

	return values
}

func (mapper *Mapper[K]) Flush() {
	mapper.maps = make(map[K]any)
}
