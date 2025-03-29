package gv1

type Auther interface {
	AddUser(string, string) (int, error)
}