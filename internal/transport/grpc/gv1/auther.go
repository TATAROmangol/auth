package gv1

type Auther interface {
	Register(string, string) (string, error)
	Login(string, string) (string, error)
}