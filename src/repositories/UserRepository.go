package repositories

import (
	"configs"
	"fmt"
	"models"
	u "utils"
)

// UserCreateFriend Create Friend Relationship
func UserCreateFriend(email1, email2 string) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}
	// Check if email1 - email2 are friends ?
	user1, _ := UserGetByEmail(email1)
	if user1.ID == 0 {
		return relationship, fmt.Errorf("User %s not exits", email1)
	}

	user2, _ := UserGetByEmail(email2)
	if user2.ID == 0 {
		return relationship, fmt.Errorf("User %s not exits", email2)
	}

	db.Where("user1_id = ? AND user2_id = ?", user1.ID, user2.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_YES
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.User1ID = user1.ID
	relationship.User2ID = user2.ID
	relationship.FriendStatus = configs.FRIEND_STATUS_YES
	relationship.Subscribe = configs.SUBSRIBE_YES
	db.Create(&relationship)

	return relationship, nil
}

// UserGetByEmail Get User by email
func UserGetByEmail(email string) (models.User, error) {
	db := u.DbConn()

	user := models.User{}
	user.ID = 0
	user.Email = ""

	db.Where("email = ?", email).First(&user)

	return user, nil
}

// UserCreate Create User
func UserCreate(email string) (models.User, error) {
	db := u.DbConn()

	user := models.User{}
	user.Email = email

	db.Create(&user)

	return user, nil
}
