package repository

type RoomFilter struct {
	IncludeDeleted bool // true/false
	// Company *string  // nil == ignore, or company name
}
