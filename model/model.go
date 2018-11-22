package model

type db interface {
	SelectPeople() ([]*User, error)
}

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) Users() ([]*User, error) {
	return m.SelectPeople()
}

type User struct {
	Id        string
	FirstName string
	LastName  string
}
