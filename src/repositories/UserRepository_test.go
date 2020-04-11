package repositories

import (
	"fmt"
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

	// Init data
	UserCreate("email@test.com")

	// Test
	user, err := UserGetByEmail("email@test.com")

	if err != nil {
		t.Errorf("error = %s; want nil", err.Error())
	}

	if user.Email != "email@test.com" {
		t.Errorf("user.email = %s; want email@test.com", user.Email)
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
