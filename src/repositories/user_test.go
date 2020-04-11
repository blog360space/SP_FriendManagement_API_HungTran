package repositories

import (
	"testing"
	u "utils"
	"configs"
	"fmt"
)

func Test_User_CreateFriend (t *testing.T) {
	u.TruncateTable("users")
	u.TruncateTable("relationship")

	email1 := "andy@example.com"
	email2 := "john@example.com"

	user1, _ := User_Create(email1)
	user2, _ := User_Create(email2)

	rs, err := User_CreateFriend(email1, email2)
	if (err != nil) {
		t.Errorf("error = %s; want nil", err.Error())
	}
	if (rs != true) {
		t.Errorf("rs = %t want true", rs)
	}

	db := u.DbConn()
	defer db.Close()
	var countRow string
	err3 := db.QueryRow(
		`SELECT COUNT(*) AS countRow
		FROM relationship
		WHERE user1_id = ?
		AND user2_id = ?
		AND subscribe = ?
		AND friend_status = ?
		LIMIT 1`, user1.Id, user2.Id, configs.SUBSRIBE_YES, configs.SUBSRIBE_YES).Scan(&countRow)

	if (err3 != nil) {
		t.Errorf("Table Relationship don't have new row. Error: %s", err3.Error())
	}

	if (countRow != "1") {
		t.Errorf("Table Relationship countRow = %s, want 1. Error: %s", countRow, err3.Error())
	}

	var strExpectError string = "";
	strExpectError = fmt.Sprintf("User %s not exits", "abc@mail.com");

	rs1, err1 := User_CreateFriend("abc@mail.com", email2)
	if (err1.Error() != strExpectError) {
		t.Errorf("error = %s; want '%s'", err1.Error(), strExpectError)
	}
	if (rs1 != false) {
		t.Errorf("rs1 = %t want false", rs1)
	}

	rs2, err2 := User_CreateFriend(email1, "abc@mail.com")
	if (err2.Error() != strExpectError) {
		t.Errorf("error = %s; want '%s'", err2.Error(), strExpectError)
	}
	if (rs2 != false) {
		t.Errorf("rs2 = %t want false", rs2)
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
