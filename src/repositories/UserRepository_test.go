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

func TestUserCreateFriendError(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	email1 := "andy@example.com"
	email2 := "john@example.com"

	UserCreate(email1)
	UserCreate(email2)

	var strExpectError string = ""
	strExpectError = fmt.Sprintf("User %s not exits", "abc@mail.com")

	_, err1 := UserCreateFriend("abc@mail.com", email2)
	if err1.Error() != strExpectError {
		t.Errorf("error = %s; want '%s'", err1.Error(), strExpectError)
	}

	_, err2 := UserCreateFriend(email1, "abc@mail.com")

	if err2.Error() != strExpectError {
		t.Errorf("error = %s; want '%s'", err2.Error(), strExpectError)
	}

}

func TestUserCreateFriendSucess(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	email1 := "andy@example.com"
	email2 := "john@example.com"

	user1, err1 := UserCreate(email1)
	user2, err2 := UserCreate(email2)

	relationship, err := UserCreateFriend(email1, email2)
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
	relationship1, _ := UserCreateFriend(email1, email2)

	if relationship1.ID != relationship.ID {
		t.Errorf("Update exits relationship error")
	}
}

func TestUserGetFriendsByEmail(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	email1 := "andy@example.com"
	email2 := "john@example.com"
	email3 := "hungtran@example.com"

	_, err1 := UserGetFriendsByEmail(email3)
	expectedMsg := fmt.Sprintf("User %s not exits", email3)
	if err1.Error() != expectedMsg {
		t.Errorf("error = %s; want %s", err1.Error(), expectedMsg)
	}

	UserCreate(email1)
	UserCreate(email2)
	UserCreate(email3)

	UserCreateFriend(email1, email2)
	UserCreateFriend(email1, email3)

	users, err := UserGetFriendsByEmail(email1)

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

	email1 := "andy@example.com"
	email2 := "john@example.com"
	email3 := "hungtran@example.com"

	UserCreate(email1)
	UserCreate(email2)
	UserCreate(email3)

	users, err = UserGetFriendsCommon(email1, email2)
	expectedMsg := fmt.Sprintf("No common friend for %s and %s", email1, email2)
	// User has no common friend.
	if expectedMsg != err.Error() {
		t.Errorf("error = %s; want '%s'", err.Error(), expectedMsg)
	}

	UserCreateFriend(email3, email1)
	UserCreateFriend(email3, email2)

	users, err = UserGetFriendsCommon(email1, email2)

	if err != nil {
		t.Errorf("error = %s; want null", err.Error())
	}

	if len(users) != 1 {
		t.Errorf("len(user) = %d incorrect; want %d", len(users), 1)
	}

	if users[0].Email != email3 {
		t.Errorf("user.Email = %s incorrect; want %s", users[0].Email, email3)
	}
}

func TestUserSubscribe(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var relationship models.Relationship

	email1 := "andy@example.com"
	email2 := "john@example.com"

	u1, _ := UserCreate(email1)
	u2, _ := UserCreate(email2)
	relationship, _ = UserSubscribe(email1, email2)

	if relationship.ID == 0 {
		t.Errorf("relationship.ID = %d incorrect; want > 0", relationship.ID)
	}

	if relationship.Subscribe != configs.SUBSRIBE_YES {
		t.Errorf("relationship.ISubscribeD = %d incorrect; want %d", relationship.ID, configs.SUBSRIBE_YES)
	}

	if relationship.User1ID != u1.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User1ID, u1.ID)

	}

	if relationship.User2ID != u2.ID {
		t.Errorf("relationship.User1ID != u1.ID (%d != %d) ; want equal", relationship.User2ID, u2.ID)

	}
}

func TestUserBlock(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var relationship models.Relationship

	email1 := "andy@example.com"
	email2 := "john@example.com"

	requestor, _ := UserCreate(email1)
	target, _ := UserCreate(email2)
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
