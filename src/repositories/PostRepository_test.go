package repositories

import (
	"models"
	"testing"
	u "utils"
)

func TestPostCreatePost(t *testing.T) {
	u.DbTruncateTable("posts")

	postText := "Hello world"
	user, _ := UserCreate("hung.tran@example.com")
	post, err := PostCreatePost(user, postText)

	if err != nil {
		t.Errorf("PostCreate() error = %s; want null", err.Error())
	}

	if post.ID == 0 {
		t.Errorf("PostCreate() post.ID = %d; want > 0", post.ID)
	}

	if post.Text != postText {
		t.Errorf("PostCreate() post.Text = '%s'; want '%s'", post.Text, postText)
	}
}

func TestPostGetPostRecipients(t *testing.T) {
	u.DbTruncateTable("posts")
	u.DbTruncateTable("users")
	u.DbTruncateTable("relationships")

	var recipients []models.User
	var err error
	var u1, u2, u3 models.User
	var post models.Post

	post, _ = PostCreatePost(u1, "Hello world")

	u1, err = UserCreate("user1@example.com")
	u2, err = UserCreate("user2@example.com")
	u3, err = UserCreate("user3@example.com")

	UserCreateFriend(u2, u1)
	UserSubscribe(u3, u1)

	recipients, err = PostGetPostRecipients(u1, post)

	if err != nil {
		t.Errorf("PostGetPostRecipients() recipients = %s; want null", err.Error())
	}

	if recipients[0].Email != u3.Email {
		t.Errorf("PostGetPostRecipients() users[0]='%s', want %s", recipients[0].Email, u3.Email)
	}

	if recipients[1].Email != u2.Email {
		t.Errorf("PostGetPostRecipients() users[1]='%s', want %s", recipients[1].Email, u2.Email)
	}
}
