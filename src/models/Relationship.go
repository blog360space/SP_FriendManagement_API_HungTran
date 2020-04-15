package models

// Relationship Struct
type Relationship struct {
	ID           uint
	RequestorID      uint
	TargetID     uint
	Subscribe    int
	FriendStatus int
}
