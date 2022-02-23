package repository

type RoomFilter struct {
	IncludeDeleted bool // true/false
	// TODO uncomment and add condition processing in the respective repository
	// Company *string  // nil == ignore, or company name
}
