package gv1

type Auther interface {
	CreateUser(string, string) (int, error)
}