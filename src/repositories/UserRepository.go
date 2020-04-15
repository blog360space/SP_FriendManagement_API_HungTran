package repositories

import (
	"configs"
	"fmt"
	"models"
	u "utils"

	"gopkg.in/go-playground/validator.v9"
)

// UserCreateFriend Create Friend Relationship
func UserCreateFriend(requestor, target models.User) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}
	// Check if email1 - email2 are friends ?
	db.Where("requestor_id = ? AND target_id = ?", requestor.ID, target.ID).First(&relationship)

	if relationship.ID > 0 {
		if relationship.FriendStatus == configs.FRIEND_STATUS_BLOCK {
			return relationship, fmt.Errorf("Target %s Blocked Requestor %s", requestor.Email, target.Email)
		}

		relationship.FriendStatus = configs.FRIEND_STATUS_YES
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.RequestorID = requestor.ID
	relationship.TargetID = target.ID
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

	validate := validator.New()
	err := validate.Var(email, "email")
	if err != nil {
		return user, fmt.Errorf("Please input a valid email address")
	}

	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return user, fmt.Errorf("User %s not exits", email)
	}

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
		SELECT r.target_id
		FROM relationships r
		WHERE r.requestor_id = ?
	)`

	db.Raw(sql, user.ID).Scan(&users)

	return users, nil
}

// UserGetFriendsCommon Get user friends common
func UserGetFriendsCommon(requestor, target models.User) ([]models.User, error) {
	db := u.DbConn()
	var users = []models.User{}

	sql := `SELECT *
	FROM users
	WHERE id IN (
		SELECT r.requestor_id
		FROM relationships r
		WHERE r.target_id IN (?, ?)
	)`

	db.Raw(sql, requestor.ID, target.ID).Scan(&users)

	if len(users) == 0 {
		return users, fmt.Errorf("No common friend for %s and %s", requestor.Email, target.Email)
	}

	return users, nil
}

// UserSubscribe subscribe to updates from an email address.
func UserSubscribe(requestor, target models.User) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}

	db.Where("requestor_id = ? AND target_id = ?", requestor.ID, target.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_YES
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.RequestorID = requestor.ID
	relationship.TargetID = target.ID
	relationship.FriendStatus = configs.FRIEND_STATUS_NO
	relationship.Subscribe = configs.SUBSRIBE_YES
	db.Create(&relationship)

	return relationship, nil
}

// UserBlock As a user, I need an API to block updates from an email address.
// {"requestor": "lisa@example.com","target": "john@example.com"}
func UserBlock(requestor, target models.User) (models.Relationship, error) {
	db := u.DbConn()
	relationship := models.Relationship{}

	db.Where("requestor_id = ? AND target_id = ?", requestor.ID, target.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_BLOCK
		relationship.Subscribe = configs.SUBSRIBE_NO
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.RequestorID = requestor.ID
	relationship.TargetID = target.ID
	relationship.FriendStatus = configs.FRIEND_STATUS_BLOCK
	relationship.Subscribe = configs.SUBSRIBE_NO
	db.Create(&relationship)

	return relationship, nil
}

// UserGetSubscribeUsers get users subscribed
func UserGetSubscribeUsers(user models.User) ([]models.User, error) {
	db := u.DbConn()

	var subscribeUsers []models.User

	sql := `SELECT *
	FROM users u
	WHERE id IN (
		SELECT r.requestor_id
		FROM relationships r
		WHERE r.target_id = ?
	)
	ORDER BY u.id DESC`

	db.Raw(sql, user.ID).Scan(&subscribeUsers)

	return subscribeUsers, nil
}

// UserRegister
func UserRegister(email string) (models.User, error) {
	var err error
	var user models.User
	// validate email
	validate := validator.New()
	err = validate.Var(email, "email")
	if err != nil {
		return user, fmt.Errorf("Please input a valid email address")
	}
	user, err = UserGetByEmail(email)
	if user.ID > 0 {
		return user, fmt.Errorf("Email %s is exits, please use other email", email)
	}

	user, err = UserCreate(email)

	return user, err
}
