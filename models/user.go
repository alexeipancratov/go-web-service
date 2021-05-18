package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers Returns slice of user pointers
func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		// We cannot return nil because we're not returning a pointer value
		return User{}, errors.New("new User must not include id or it must be 0")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range GetUsers() {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
