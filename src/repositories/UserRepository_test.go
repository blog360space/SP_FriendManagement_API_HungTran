package repositories

import (
	"configs"
	"fmt"
	"models"
	"testing"
	u "utils"
)

func TestUserCreate(t *testing.T) {
	u.DbTruncateTable("users")

	user, err := UserCreate("email@test.com")

	if err != nil {
		t.Errorf("err = %s; want nil", err.Error())
	}

	if user.ID == 0 {
		t.Errorf("user.id = %d; want > 0", user.ID)
	}

	if user.Email != "email@test.com" {
		t.Errorf("user.email = %s; want email@test.com", user.Email)
	}
}

func TestUserGetByEmail(t *testing.T) {
	u.DbTruncateTable("users")
	var user models.User
	var err error

	// Init data
	UserCreate("email@test.com")

	// Test
	user, err = UserGetByEmail("email@test.com")

	if err != nil {
		t.Errorf("error = %s; want nil", err.Error())
	}

	if user.Email != "email@test.com" {
		t.Errorf("user.email = %s; want email@test.com", user.Email)
	}

	// Test
	user, err = UserGetByEmail("email1@test.com")
	var exptectedMsg string = fmt.Sprintf("User %s not exits", "email1@test.com")
	if err.Error() != exptectedMsg {
		t.Errorf("error = %s, want '%s'", err.Error(), exptectedMsg)
	}
}

func TestUserCreateFriendSucess(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	user1, err1 := UserCreate("andy@example.com")
	user2, err2 := UserCreate("john@example.com")

	relationship, err := UserCreateFriend(user1, user2)
	if err != nil {
		t.Errorf("error = %s; want nil", err.Error())
	}

	if relationship.User1ID != user1.ID || relationship.User2ID != user2.ID {
		t.Errorf("User1ID != user1.ID OR relationship.User2ID != user2.ID")
	}

	if err1 != nil {
		t.Errorf("error = %s; want null", err1.Error())
	}

	if err2 != nil {
		t.Errorf("error = %s; want null", err1.Error())
	}

	// Update relationship
	relationship1, _ := UserCreateFriend(user1, user2)

	if relationship1.ID != relationship.ID {
		t.Errorf("Update exits relationship error")
	}
}

func TestUserGetFriendsByEmail(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")
	var err error
	var u1, u2, u3 models.User

	_, err1 := UserGetFriendsByEmail("hungtran@example.com")
	expectedMsg := fmt.Sprintf("User %s not exits", "hungtran@example.com")
	if err1.Error() != expectedMsg {
		t.Errorf("error = %s; want %s", err1.Error(), expectedMsg)
	}

	u1, err = UserCreate("andy@example.com")
	u2, err = UserCreate("john@example.com")
	u3, err = UserCreate("hungtran@example.com")

	UserCreateFriend(u1, u2)
	UserCreateFriend(u1, u3)

	users, err := UserGetFriendsByEmail(u1.Email)

	if err != nil {
		t.Errorf("error = %s; want null", err.Error())
	}

	if len(users) != 2 {
		t.Errorf("len(user) = %d incorrect; want %d", len(users), 2)
	}
}

func TestUserGetFriendsCommon(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")
	var err error
	var users []models.User
	var u1, u2, u3 models.User

	u1, err = UserCreate("andy@example.com")
	u2, err = UserCreate("john@example.com")
	u3, err = UserCreate("hungtran@example.com")

	users, err = UserGetFriendsCommon(u1, u2)
	expectedMsg := fmt.Sprintf("No common friend for %s and %s", u1.Email, u2.Email)
	// User has no common friend.
	if expectedMsg != err.Error() {
		t.Errorf("error = %s; want '%s'", err.Error(), expectedMsg)
	}

	UserCreateFriend(u3, u1)
	UserCreateFriend(u3, u1)

	users, err = UserGetFriendsCommon(u1, u2)

	if err != nil {
		t.Errorf("error = %s; want null", err.Error())
	}

	if len(users) != 1 {
		t.Errorf("len(user) = %d incorrect; want %d", len(users), 1)
	}

	if users[0].Email != u3.Email {
		t.Errorf("user.Email = %s incorrect; want %s", users[0].Email, u3.Email)
	}
}

func TestUserSubscribe(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")
	var relationship models.Relationship

	requestor, _ := UserCreate("andy@example.com")
	target, _ := UserCreate("john@example.com")
	relationship, _ = UserSubscribe(requestor, target)

	if relationship.ID == 0 {
		t.Errorf("relationship.ID = %d incorrect; want > 0", relationship.ID)
	}

	if relationship.Subscribe != configs.SUBSRIBE_YES {
		t.Errorf("relationship.ISubscribeD = %d incorrect; want %d", relationship.ID, configs.SUBSRIBE_YES)
	}

	if relationship.User1ID != requestor.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User1ID, requestor.ID)

	}

	if relationship.User2ID != target.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User2ID, target.ID)

	}
}

func TestUserBlock(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var relationship models.Relationship

	requestor, _ := UserCreate("andy@example.com")
	target, _ := UserCreate("john@example.com")
	relationship, _ = UserBlock(requestor, target)

	if relationship.ID == 0 {
		t.Errorf("relationship.ID = %d incorrect; want > 0", relationship.ID)
	}

	if relationship.Subscribe != configs.SUBSRIBE_NO {
		t.Errorf("relationship.ISubscribeD = %d incorrect; want %d", relationship.ID, configs.SUBSRIBE_NO)
	}

	if relationship.User1ID != requestor.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User1ID, requestor.ID)

	}

	if relationship.User2ID != target.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User2ID, target.ID)

	}
}

func TestUserGetSubscribeUser(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var u1, u2, u3 models.User

	u1, _ = UserCreate("user1@example.com")
	u2, _ = UserCreate("user2@example.com")
	u3, _ = UserCreate("user3@example.com")

	UserSubscribe(u2, u1)
	UserSubscribe(u3, u1)

	users, err := UserGetSubscribeUsers(u1)

	if err != nil {
		t.Errorf("UserGetSubscribeUsers() err='%s', want nil", err.Error())
	}

	countSubscribe := len(users)
	if countSubscribe != 2 {
		t.Errorf("UserGetSubscribeUsers() countSubscribe='%d', want %d", countSubscribe, 2)
	}

	if users[0].Email != u3.Email {
		t.Errorf("UserGetSubscribeUsers() users[0]='%s', want %s", users[0].Email, u3.Email)
	}

	if users[1].Email != u2.Email {
		t.Errorf("UserGetSubscribeUsers() users[1]='%s', want %s", users[1].Email, u2.Email)
	}
}
