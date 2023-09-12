package contacts

type Nullable interface {

	// IsEmpty Determine if the map or collections is empty.
	IsEmpty() bool

	// IsNotEmpty Determine if the map or collections is not empty.
	IsNotEmpty() bool
}
