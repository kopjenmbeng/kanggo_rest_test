package dto

type Customer struct {
	Id             string
	Email          string
	FullName       string
	Salt           string
	Password       string
	Iteration      int
	SecurityLength int
}
