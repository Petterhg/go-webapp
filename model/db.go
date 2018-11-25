package model

type db interface {
	SelectUser() ([]*User, error)
}
