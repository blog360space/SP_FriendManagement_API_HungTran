package repositories

import (
	"configs"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"models"
	u "utils"
)

// UserCreateFriend Create Friend Relationship
func UserCreateFriend(user1, user2 models.User) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}
	// Check if email1 - email2 are friends ?
	db.Where("user1_id = ? AND user2_id = ?", user1.ID, user2.ID).First(&relationship)

	if relationship.ID > 0 {
		if relationship.FriendStatus == configs.FRIEND_STATUS_BLOCK {
			return relationship, fmt.Errorf("Target %s Blocked Requestor %s", user1.Email, user2.Email)
		}

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

	validate := validator.New();
	err := validate.Var(email, "email")
	if (err != nil) {
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
		SELECT r.user2_id
		FROM relationships r
		WHERE r.user1_id = ?
	)`

	db.Raw(sql, user.ID).Scan(&users)

	return users, nil
}

// UserGetFriendsCommon Get user friends common
func UserGetFriendsCommon(user1, user2 models.User) ([]models.User, error) {
	db := u.DbConn()
	var users = []models.User{}

	sql := `SELECT *
	FROM users
	WHERE id IN (
		SELECT r.user1_id
		FROM relationships r
		WHERE r.user2_id IN (?, ?)
	)`

	db.Raw(sql, user1.ID, user2.ID).Scan(&users)

	if len(users) == 0 {
		return users, fmt.Errorf("No common friend for %s and %s", user1.Email, user2.Email)
	}

	return users, nil
}

// UserSubscribe subscribe to updates from an email address.
func UserSubscribe(requestor, target models.User) (models.Relationship, error) {
	db := u.DbConn()

	relationship := models.Relationship{}

	db.Where("user1_id = ? AND user2_id = ?", requestor.ID, target.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_YES
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.User1ID = requestor.ID
	relationship.User2ID = target.ID
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

	db.Where("user1_id = ? AND user2_id = ?", requestor.ID, target.ID).First(&relationship)

	if relationship.ID > 0 {
		relationship.FriendStatus = configs.FRIEND_STATUS_BLOCK
		relationship.Subscribe = configs.SUBSRIBE_NO
		db.Save(&relationship)
		return relationship, nil
	}

	relationship.User1ID = requestor.ID
	relationship.User2ID = target.ID
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
		SELECT r.user1_id
		FROM relationships r
		WHERE r.user2_id = ?
	)
	ORDER BY u.id DESC`

	db.Raw(sql, user.ID).Scan(&subscribeUsers)

	return subscribeUsers, nil
}


// UserRegister
func UserRegister(email string)  (models.User, error) {
	var err error
	var user models.User
	// validate email
	validate := validator.New();
	err = validate.Var(email, "email")
	if (err != nil) {
		return user, fmt.Errorf("Please input a valid email address")
	}
	user, err = UserGetByEmail(email)
	if (user.ID > 0) {
		return user, fmt.Errorf( "Email %s is exits, please use other email", email)
	}

	user, err = UserCreate(email)

	return user, err
}
