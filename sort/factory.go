package sort

import (
	"github.com/goextension/mapper/contacts/sort"
	"github.com/goextension/mapper/sort/observer"
)

func NewSorter[K string | int, V any]() sort.Sorter[K, V] {
	return &MapSorter[K, V]{
		observer: &observer.SortEventObserver{},
	}
}
