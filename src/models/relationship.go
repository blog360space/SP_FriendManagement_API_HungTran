package models

type Relationship struct {
	User1Id int `json:"user1id"`
	User2Id int `json:"user2id"`
	Subscribe int `json:"subscribe"`
	FriendStatus int `json:"friend_status"`
}
