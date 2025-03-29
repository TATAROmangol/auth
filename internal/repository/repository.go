package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct{
	db *sql.DB
}

func New(db *sql.DB) *Repository{
	return &Repository{db}
}

func (r *Repository) CheckLogin(login string) (bool, error){
	stmt, err := r.db.Prepare(`
		SELECT EXIST(
			SELECT 1
			FROM users
			WHERE login = ($1)
		)
	`)
	if err != nil{
		return true, fmt.Errorf("failed in check login: %v", err)
	}

	var exist bool 
	if err := stmt.QueryRow(login).Scan(&exist); err != nil{
		return true, fmt.Errorf("failed in check login: %v", err)
	}
	return exist, nil
}

func (r *Repository) CreateUser(login, password string) (int, error){
	id := -1

	stmt, err := r.db.Prepare(`
		INSERT INTO users (login, password)
		VALUES ($1, $2) 
		RETURNING id
	`)
	if err != nil {
		return id, fmt.Errorf("failed create user in data base: %v", err)
	}
	defer stmt.Close()

	if err := stmt.QueryRow(login, password).Scan(&id); err != nil{
		return id, err
	}

	return id, nil
}

func (r *Repository) CheckUser(login, password string) (int, error){
	id := -1

	stmt, err := r.db.Prepare(`
		SELECT id
		FROM users
		WHERE login = ($1) AND password = ($2)
	`)
	if err != nil{
		return id, fmt.Errorf("failed in check user: %v", err)
	}

	if err := stmt.QueryRow(login, password).Scan(&id); err != nil{
		return id, fmt.Errorf("failed in check user: %v", err)
	}
	if id == -1{
		return id, fmt.Errorf("incorrect user data: %v", err)
	}

	return id, nil
}