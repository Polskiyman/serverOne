package store

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"

	"service/internal"
	"service/internal/user"
)

type StoreInterface interface {
	CreateUser(u user.User) (int, error)
	GetAll() ([]user.User, error)
	GetUserById(id int) (user.User, error)
	DeleteUser(id int) error
	UpdateUserFriends(u user.User) error
	UpdateUserAge(u user.User) error
}

type Store struct {
	db *sql.DB
}

func NewStore(config internal.DbConfig) *Store {
	db := connectDb(config)
	return &Store{
		db: db,
	}
}

func connectDb(config internal.DbConfig) (db *sql.DB) {
	cfg := mysql.Config{
		User:   config.User,
		Passwd: config.Passwd,
		Net:    "tcp",
		Addr:   config.Addr,
		DBName: config.DbName,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return
}

func (s *Store) CreateUser(u user.User) (int, error) {
	result, err := s.db.Exec("INSERT INTO user (name, age) VALUES (?, ?)", u.Name, u.Age)
	if err != nil {
		return 0, fmt.Errorf("put user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("put user: %v", err)
	}
	return int(id), nil
}

func (s *Store) MakeFriends(id1, id2 int) error {
	_, err := s.db.Exec("UPDATE user SET friends = ? WHERE id = ?", strconv.Itoa(id2), id1)
	if err != nil {
		return fmt.Errorf("make friends: %v", err)
	}

	_, err = s.db.Exec("UPDATE user SET friends = ? WHERE id = ?", strconv.Itoa(id1), id2)
	if err != nil {
		return fmt.Errorf("make friends: %v", err)
	}

	return nil
}

func (s *Store) GetAll() ([]user.User, error) {
	var res []user.User

	rows, err := s.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			u       user.User
			friends string
		)

		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &friends); err != nil {
			return nil, fmt.Errorf("get all users: %v", err)
		}

		for _, v := range strings.Split(friends, ",") {
			f, err := strconv.Atoi(v)
			if err != nil {
				log.Println(err)
				continue
			}
			u.Friends = append(u.Friends, f)
		}
		res = append(res, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	return res, nil
}

func (s *Store) GetUserById(id int) (user.User, error) {
	var u user.User
	var friends string
	row := s.db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	if err := row.Scan(&u.Id, &u.Name, &u.Age, &friends); err != nil {
		return user.User{}, err
	}
	for _, v := range strings.Split(friends, ",") {
		f, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			continue
		}
		u.Friends = append(u.Friends, f)
	}
	return u, nil
}

func (s *Store) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete user: %v", err)
	}
	return nil
}

func (s *Store) UpdateUserFriends(u user.User) error {
	friends := convertFriends(u)
	_, err := s.db.Exec("UPDATE user SET friends = ? WHERE id = ?", friends, u.Id)
	if err != nil {
		return fmt.Errorf("update user friends: %v", err)
	}
	return nil
}

func (s *Store) UpdateUserAge(u user.User) error {
	_, err := s.db.Exec("UPDATE user SET Age = ? WHERE id = ?", u.Age, u.Id)
	if err != nil {
		return fmt.Errorf("update user age: %v", err)
	}
	return nil
}

func convertFriends(u user.User) string {
	sFriends := make([]string, 0, len(u.Friends))
	for _, f := range u.Friends {
		sf := strconv.Itoa(f)
		sFriends = append(sFriends, sf)
	}
	return strings.Join(sFriends, ",")
}
