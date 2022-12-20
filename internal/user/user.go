package user

import "fmt"

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func (u *User) ToString() string {
	return fmt.Sprintf("Name is %s , Age %d is , friends %d and Id:%d. ", u.Name, u.Age, u.Friends, u.Id)
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetId(i int) {
	u.Id = i
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func (u *User) GetFriends() []int {
	return u.Friends
}

func (u *User) AddFriend(id int) {
	u.Friends = append(u.Friends, id)
}

func (u *User) DeleteFriend(id int) {
	for i := range u.Friends {
		if u.Friends[i] == id {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:len(u.Friends)]...)
			return
		}
	}
	return
}
