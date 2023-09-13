package contacts

type Mappable[K string | int, V any] interface {
	Has(haystack K) bool

	SetMap(haystack K, needle V) Mappable[K, V]

	Equal(haystack K, closure func(value V) bool) bool

	Get(haystack K) V

	Unset(haystack K)

	Keys() []K

	Values() []V

	Enumerable[K, V]

	Nullable

	Clearable
}
