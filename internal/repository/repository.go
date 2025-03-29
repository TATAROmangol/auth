package repository

import "context"

type Repository struct{
	ctx context.Context
	
}

func CreateUser(login, password string) (int, error){
	return 0, nil
}

func CheckUser(login, password string){

}