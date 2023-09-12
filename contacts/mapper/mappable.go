package mapper

import "goextension/core/contacts"

type Mappable[K string | int, V any] interface {
	Has(haystack K) bool

	SetMap(haystack K, needle V) Mappable[K, V]

	Get(haystack K) V

	Unset(haystack K)

	Keys() []K

	Values() []V

	contacts.Clearable
}
