package service

type Mock struct{}

func (Mock) CreateUser(name string, age int) (int, error) {
	if name == "Biba" {
		return 1, nil
	}
	if name == "Boba" {
		return 2, nil
	}
	return 0, nil
}

func (Mock) GetAllUsers() (string, error) {
	return "all users", nil
}

func (Mock) MakeFriends(id1, id2 int) (name1, name2 string, err error) {
	return "Biba", "Boba", nil
}

func (Mock) DeleteUser(id int) (string, error) {
	return "Biba", nil
}

func (Mock) GetUserFriends(id int) (res string, err error) {
	return "Boba", nil
}

func (Mock) UpdateAge(id, age int) error {
	return nil
}
