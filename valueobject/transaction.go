package valueobject

import "github.com/google/uuid"

// Transctions is a valueobject because it has no identifier and unmutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt uuid.Time
}
