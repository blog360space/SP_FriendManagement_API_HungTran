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

// UserGetFriendsByEmail Get Friends by email.
func UserGetFriendsByEmail(email string) ([]models.User, error) {
	db := u.DbConn()

	var users = []models.User{}
	user, _ := UserGetByEmail(email)
	if user.ID == 0 {
		return users, fmt.Errorf("User %s not exits", email)
	}

	sql := `SELECT *
	FROM users
	WHERE id IN (
		SELECT r.user2_id
		FROM relationships r
		WHERE r.user1_id = ?
	)`

	db.Raw(sql, user.ID).Scan(&users)

	return users, nil
}

// UserGetFriendsCommon Get user friends common
func UserGetFriendsCommon(email1, email2 string) ([]models.User, error) {
	db := u.DbConn()
	var users = []models.User{}

	user1, _ := UserGetByEmail(email1)
	if user1.ID == 0 {
		return users, fmt.Errorf("User %s not exits", email1)
	}

	user2, _ := UserGetByEmail(email2)
	if user2.ID == 0 {
		return users, fmt.Errorf("User %s not exits", email2)
	}

	sql := `SELECT *
	FROM users
	WHERE id IN (
		SELECT r.user1_id
		FROM relationships r
		WHERE r.user2_id IN (?, ?)
	)`

	db.Raw(sql, user1.ID, user2.ID).Scan(&users)

	if len(users) == 0 {
		return users, fmt.Errorf("No common friend for %s and %s", email1, email2)
	}

	return users, nil
}

// UserSubscribe subscribe to updates from an email address.
func UserSubscribe(requestor, target string) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}
	// Check if email1 - email2 are friends ?
	user1, _ := UserGetByEmail(requestor)
	if user1.ID == 0 {
		return relationship, fmt.Errorf("Requestor %s not exits", requestor)
	}

	user2, _ := UserGetByEmail(target)
	if user2.ID == 0 {
		return relationship, fmt.Errorf("Target %s not exits", target)
	}

	db.Where("user1_id = ? AND user2_id = ?", user1.ID, user2.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_YES
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.User1ID = user1.ID
	relationship.User2ID = user2.ID
	relationship.FriendStatus = configs.FRIEND_STATUS_NO
	relationship.Subscribe = configs.SUBSRIBE_YES
	db.Create(&relationship)

	return relationship, nil
}
