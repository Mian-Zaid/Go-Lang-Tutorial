package controllers

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	name string
	id   int64
}

func (user *User) InitUser() {
	print("Enter user name: ")
	fmt.Scan(&user.name)

	print("Enter user Id: ")
	fmt.Scan(&user.id)
}

func (user *User) PrintUser() {
	println("Name => ", user.name)
	println("ID => ", user.id)
}

func (user *User) GenerateRandomUser() User {
	// List of example names
	names := []string{"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace"}

	// Initialize the random number generator using the current time as the seed
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a name from the list
	randomIndex := rand.Intn(len(names))

	// Generate a random ID using the rand.Int63() function
	randomID := rand.Int63()

	// Create and return the random User
	return User{
		name: names[randomIndex],
		id:   randomID,
	}
}
