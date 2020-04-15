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

	_, err = UserGetByEmail("blabla")
	if (err.Error() != "Please input a valid email address") {
		t.Errorf("error = %s; want 'Please input a valid email address'", err.Error())
	}

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

	requestor, err1 := UserCreate("andy@example.com")
	target, err2 := UserCreate("john@example.com")

	relationship, err := UserCreateFriend(requestor, target)
	if err != nil {
		t.Errorf("error = %s; want nil", err.Error())
	}

	if relationship.RequestorID != requestor.ID || relationship.TargetID != target.ID {
		t.Errorf("requestorID != requestor.ID OR relationship.TargetID != target.ID")
	}

	if err1 != nil {
		t.Errorf("error = %s; want null", err1.Error())
	}

	if err2 != nil {
		t.Errorf("error = %s; want null", err1.Error())
	}

	// Update relationship
	relationship1, _ := UserCreateFriend(requestor, target)

	if relationship1.ID != relationship.ID {
		t.Errorf("Update exits relationship error")
	}
}

func TestUserGetFriendsByEmail(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")
	var err error
	var u1, u2, u3, ua, ub models.User

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

	ua , err = UserCreate("a@example.com")
	ub , err = UserCreate("b@example.com")

	UserCreateFriend(ua, ub)

	users1, err := UserGetFriendsByEmail(ua.Email)
	if len(users1) != 1 {
		t.Errorf("len(user) = %d incorrect; want %d", len(users), 1)
	}
	if users1[0].Email != ub.Email {
		t.Errorf("users1[0].Email = %s incorrect; want %s", users1[0].Email, ub.Email)
	}

	users2, err := UserGetFriendsByEmail(ub.Email)
	if users2[0].Email != ua.Email {
		t.Errorf("users1[0].Email = %s incorrect; want %s", users1[0].Email, ua.Email)
	}
}

func TestUserGetFriendsCommon(t *testing.T) {
	db := u.DbConn()
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

	r1 := models.Relationship{RequestorID: u1.ID, TargetID: u3.ID, FriendStatus: configs.FRIEND_STATUS_YES}
	db.Save(&r1);
	r2 := models.Relationship{RequestorID: u2.ID, TargetID: u3.ID, FriendStatus: configs.FRIEND_STATUS_YES}
	db.Save(&r2);

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

	users, err = UserGetFriendsCommon(u2, u1)
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

	if relationship.RequestorID != requestor.ID {
		t.Errorf("relationship.RequestorID != u1.ID (%d != %d) ; want equal", relationship.RequestorID, requestor.ID)

	}

	if relationship.TargetID != target.ID {
		t.Errorf("relationship.RequestorID != u1.ID (%d != %d) ; want equal", relationship.TargetID, target.ID)

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

	if relationship.RequestorID != requestor.ID {
		t.Errorf("relationship.RequestorID != u1.ID (%d != %d) ; want equal", relationship.RequestorID, requestor.ID)

	}

	if relationship.TargetID != target.ID {
		t.Errorf("relationship.RequestorID != u1.ID (%d != %d) ; want equal", relationship.TargetID, target.ID)

	}
}

func TestUserGetSubscribeUser(t *testing.T) {
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var u1, u2, u3 models.User

	u1, _ = UserCreate("requestor@example.com")
	u2, _ = UserCreate("target@example.com")
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

func TestUserRegister(t *testing.T)  {
	u.DbTruncateTable("users")
	var err error
	var u2 models.User
	var expectMsg string

	expectMsg = "Please input a valid email address"
	_, err = UserRegister("invalid email")

	if err!= nil && err.Error() != expectMsg {
		t.Errorf("UserRegister() err='%s', want'%s'", err.Error(), expectMsg)
	}

	UserCreate("requestor@example.com");
	_, err = UserRegister("requestor@example.com");
	expectMsg = fmt.Sprintf("Email %s is exits, please use other email", "requestor@example.com")

	if err!= nil && err.Error() != expectMsg {
		t.Errorf("UserRegister() err='%s', want'%s'", err.Error(), expectMsg)
	}

	u2, err = UserRegister("target@example.com");
	if err != nil {
		t.Errorf("UserRegister() err = %s; want nil", err.Error())
	}

	if u2.ID == 0 {
		t.Errorf("UserRegister() u2.ID = %d; want > 0", u2.ID)
	}

	if u2.Email != "target@example.com" {
		t.Errorf("UserRegister() u2.Email = %s; want %s", u2.Email, "target@example.com")
	}
}