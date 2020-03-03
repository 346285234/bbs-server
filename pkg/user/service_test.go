package user

import (
	"fmt"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ids := []uint{1, 2}
	users, _ := GetUsers(ids)
	fmt.Println(users)
}

func TestGetUser(t *testing.T) {
	var id uint = 1
	user, _ := GetUser(id)
	fmt.Println(user)
}
