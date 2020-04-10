package repositories

import (
	"testing"
	u "utils"
)

func Test_User_CreateFriend (t *testing.T) {
	u.TruncateTable("users")


	email1 := "andy@example.com"
	email2 := "john@example.com"

	User_Create(email1)
	User_Create(email2)

	rs, err := User_CreateFriend(email1, email2)
	if (err != nil) {
		t.Errorf("error = %s; want nil", err.Error())
	}
	if (rs != true) {
		t.Errorf("rs = %t want true", rs)
	}

	rs1, err1 := User_CreateFriend("abc@mail.com", email2)
	if (err1.Error() != "User 1 not exits") {
		t.Errorf("error = %s; want 'User 1 not exits'", err1.Error())
	}
	if (rs1 != false) {
		t.Errorf("rs1 = %t want false", rs)
	}

	rs2, err2 := User_CreateFriend(email1, "abc@mail.com")
	if (err2.Error() != "User 2 not exits") {
		t.Errorf("error = %s; want 'User 2 not exits'", err2.Error())
	}
	if (rs2 != false) {
		t.Errorf("rs2 = %t want false", rs)
	}
}

func Test_User_Create(t *testing.T)  {
	u.TruncateTable("users")

	user, err := User_Create("email@test.com")


	if (err != nil) {
		t.Errorf("err = %s; want nil", err.Error())
	}

	if (user.Id == 0) {
		t.Errorf("user.id = %d; want > 0", user.Id)
	}

	if (user.Email != "email@test.com") {
		t.Errorf("user.email = %s; want email@test.com", user.Email)
	}
}

func Test_User_GetByEmail(t *testing.T)  {
	u.TruncateTable("users")

	// Init data
	User_Create("email@test.com")

	// Test
	user, err := User_GetByEmail("email@test.com")

	if (err != nil) {
		t.Errorf("error = %s; want nil", err.Error())
	}

	if (user.Email != "email@test.com") {
		t.Errorf("user.email = %s; want email@test.com", user.Email)
	}
}
