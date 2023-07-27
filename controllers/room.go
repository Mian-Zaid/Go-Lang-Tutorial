package controllers

type Room struct {
	roomName string
	roomId   int64
	users    []User
}

func (room *Room) InitRoom(name string, id int16) {
	room.roomName = name
	room.roomId = int64(id)
	room.users = []User{}
}

func (room *Room) AddUserToRoom(user *User) {
	room.users = append(room.users, *user)
}

func (room *Room) PrintRoom() {
	println("----------Room--------")
	println("Room name => ", room.roomName)
	println("Room id => ", room.roomId)
	// println("Users => ")
	for _, user := range room.users {
		println("")
		user.PrintUser()
	}
}
