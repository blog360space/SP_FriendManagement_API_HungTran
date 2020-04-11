package repositories

import (
	u "utils"
	"models"
	"errors"
	"configs"
	"fmt"

)

func User_CreateFriend(email1, email2 string) (bool, error) {
	// Check if email1 - email2 are friends ?
	user1, _ := User_GetByEmail(email1)
	if (user1.Id == 0) {
		return false, errors.New(fmt.Sprintf("User %s not exits", email1))
	}

	user2, _ := User_GetByEmail(email2)
	if (user2.Id == 0) {
		return false, errors.New(fmt.Sprintf("User %s not exits", email2))
	}

	db := u.DbConn()
	defer db.Close()

	_, err := db.Query(
		"INSERT INTO relationship(user1_id, user2_id, subscribe, friend_status) VALUES(?, ?, ?, ?)",
		user1.Id, user2.Id, configs.SUBSRIBE_YES, configs.FRIEND_STATUS_YES)

	if (err != nil) {
		return false, err
	}

	return true, nil
}

func User_GetByEmail(email string) (models.User, error) {
	db := u.DbConn()
	defer db.Close()

	user := models.User{0, ""}

	selDb, err := db.Query("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	if (err != nil) {
		return user, err
	}

	for selDb.Next() {
		var id int
		var email string
		err = selDb.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Email = email
	}

	return user, nil
}

func User_Create(email string) (models.User, error){
	db := u.DbConn()
	defer db.Close()

	user := models.User{0, ""}

	_, err := db.Query("INSERT INTO users(email) VALUES(?)", email)
	if (err != nil) {
		return user, err
	}

	selDb, err := db.Query("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	if (err != nil) {
		return user, err
	}

	for selDb.Next() {
		var id int
		var email string
		err = selDb.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Email = email
	}

	return user, nil
}