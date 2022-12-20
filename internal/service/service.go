package service

import (
	"log"

	"service/internal"
	"service/internal/store"
	"service/internal/user"
)

type ServiceInterface interface {
	CreateUser(name string, age int) (int, error)
	GetAllUsers() (string, error)
	MakeFriends(id1, id2 int) (name1, name2 string, err error)
	DeleteUser(id int) (string, error)
	GetUserFriends(id int) (res string, err error)
	UpdateAge(id, age int) error
}

type Service struct {
	store store.StoreInterface
}

func NewService(config internal.DbConfig) *Service {
	return &Service{
		store: store.NewStore(config),
	}
}

func (s *Service) CreateUser(name string, age int) (int, error) {
	u := user.NewUser(name, age)
	return s.store.CreateUser(u)
}

func (s *Service) GetAllUsers() (string, error) {
	users, err := s.store.GetAll()
	if err != nil {
		return "", err
	}

	var res string
	for _, u := range users {
		res += u.ToString()
	}
	return res, nil
}

func (s *Service) MakeFriends(id1, id2 int) (name1, name2 string, err error) {
	user1, err := s.store.GetUserById(id1)
	if err != nil {
		return
	}

	user2, err := s.store.GetUserById(id2)
	if err != nil {
		return
	}

	user1.AddFriend(id2)
	user2.AddFriend(id1)

	err = s.store.UpdateUserFriends(user1)
	if err != nil {
		return
	}

	err = s.store.UpdateUserFriends(user2)
	if err != nil {
		return
	}

	return user1.GetName(), user2.GetName(), nil

}

func (s *Service) DeleteUser(id int) (string, error) {
	u, err := s.store.GetUserById(id)
	if err != nil {
		return "", err
	}
	err = s.store.DeleteUser(id)
	if err != nil {
		return "", err
	}

	for _, fid := range u.GetFriends() {
		u, err := s.store.GetUserById(fid)
		if err != nil {
			log.Printf("can't get friend id=%d of user with id=%d", fid, id)
			continue
		}
		u.DeleteFriend(id)
		err = s.store.UpdateUserFriends(u)
		if err != nil {
			return "", err
		}
	}

	return u.GetName(), nil
}

func (s *Service) GetUserFriends(id int) (res string, err error) {
	u, err := s.store.GetUserById(id)
	if err != nil {
		return
	}

	for _, fid := range u.GetFriends() {
		u, err := s.store.GetUserById(fid)
		if err != nil {
			log.Printf("can't get friend id=%d of user with id=%d", fid, id)
			continue
		}
		res += u.ToString()
	}
	return
}

func (s *Service) UpdateAge(id, age int) error {
	u, err := s.store.GetUserById(id)
	if err != nil {
		return err
	}
	u.SetAge(age)
	return s.store.UpdateUserAge(u)
}
