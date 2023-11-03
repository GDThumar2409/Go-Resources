package mockdb

import "fmt"

type User struct {
	ID    int
	First string
}

type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	u, ok := md.Users[id]

	if ok {
		return u, nil
	}

	return User{}, fmt.Errorf("User not found id %v", id)
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]

	if !ok {
		md.Users[u.ID] = u
		return nil
	}

	return fmt.Errorf("User already found id %v", u.ID)
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	Ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.Ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.Ds.SaveUser(u)
}
