package contacts

type Mappable[K string | int] interface {
	HasKey(haystack K) bool

	SetMap(haystack K, needle any) Mappable[K]

	Get(haystack K) any

	Delete(haystack K)

	Keys() []K

	Values() []any

	Clearable
}
