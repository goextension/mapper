package mapper

import "goextension/mapper/contacts"

func MakeMapper[K string | int, V any]() contacts.Mappable[K, V] {
	return &Mapper[K, V]{
		maps: make(map[K]V),
	}
}
