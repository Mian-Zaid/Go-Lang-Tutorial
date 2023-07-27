package main

import (
	controllers "gotutorial.com/main/controllers"
)

func main() {
	room := controllers.Room{}
	room.InitRoom("RoomTest", 1)

	for i := 0; i < 5; i++ {
		user := controllers.User{}
		user = user.GenerateRandomUser()
		room.AddUserToRoom(&user)
	}

	room.PrintRoom()

}
