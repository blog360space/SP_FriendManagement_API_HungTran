package repositories

import (
	"models"
	"strings"
	"utils"
)

// PostGetPostRecipients get post recipients
func PostGetPostRecipients(sender models.User, post models.Post) ([]models.User, error) {
	var recipients, subscribers []models.User
	var err error

	recipients, err = UserGetFriendsByEmail(sender.Email)
	if err != nil {
		return recipients, err
	}

	subscribers, err = UserGetSubscribeUsers(sender)
	if err != nil {
		return subscribers, err
	}

	recipients = append(subscribers)
	arrStr := strings.Split(post.Text, " ")
	countStr := len(arrStr)
	var userTmp models.User
	for i := 0; i < countStr; i++ {
		if utils.ValidateIsEmail(arrStr[i]) == false {
			continue
		}

		userTmp, _ = UserGetByEmail(arrStr[i])
		if userTmp.ID == 0 {
			// @todo errrmsg here
			continue
		}
		recipients = append(recipients, userTmp)
	}

	return recipients, nil
}

// PostCreatePost Create new post
func PostCreatePost(user models.User, postText string) (models.Post, error) {
	db := utils.DbConn()

	var post models.Post
	post.Text = postText
	post.UserID = user.ID
	db.Create(&post)

	return post, nil
}
