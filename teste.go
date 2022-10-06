package main

import (
	"database/sql"
	"fmt"
	"jwt-authentication-golang/model"
	"log"
  "golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "goauth"
	USER     = "admin"
	PASSWORD = "admin"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInterface interface {
	HashPassword(password string) error
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func main() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("Could not connect to Database")
		log.Fatal("error", err)
	} else {
		fmt.Println("Connected")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO users values($1, $2, $3, $4, $5)"

	insert, err := db.Prepare(sql)
	checkErr(err)

	user := &User{
		ID:       2,
		Name:     "Lenilson",
		Email:    "lenilsonts@gmail.com",
		Username: "lenilsonts",
		Password: "spring",
	}

	user.HashPassword(user.Password)

	result, err := insert.Exec(user.ID, user.Name, user.Email, user.Username, user.Password)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	defer db.Close()
}
