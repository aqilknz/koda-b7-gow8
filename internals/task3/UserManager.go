package task3

import "fmt"

type User struct {
	Id   int
	Name string
}

type UserManager struct {
	users map[int]*User
}

func NewUserManager() *UserManager {
	new := make(map[int]*User)
	return &UserManager{
		users: new,
	}
}
func (um *UserManager) AddUser(id int, nama string) {
	if _, orang := um.users[id]; orang {
		fmt.Println("user sudah tersedia")
		return
	}
	um.users[id] = &User{Id: id, Name: nama}
}
func (um *UserManager) GetUser(id int) string {
	if um.users[id] == nil {
		return "user tidak ditemukan"
	}
	user := um.users[id]
	return fmt.Sprintf("ID: %d\nNama: %s\n", user.Id, user.Name)

}
