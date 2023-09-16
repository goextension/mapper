package maps

import (
	"github.com/goextension/contacts/collection"
	"github.com/goextension/mapper/contacts/sort"
)

type Mappable[K string | int, V any] interface {
	Has(haystack K) bool

	Store(haystack K, needle V) Mappable[K, V]

	Equal(haystack K, closure func(value V) bool) bool

	Get(haystack K) V

	Unset(haystack K)

	DeleteWithCopy(haystack K) V

	Keys() []K

	Values() []V

	Enumerable[K, V]

	sort.Sortable[K, V]

	collection.Nullable

	collection.Flusher

	collection.Countable
}
