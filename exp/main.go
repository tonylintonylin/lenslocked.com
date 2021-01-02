package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	us.DestructiveReset()

	user := models.User{
		Name:  "Mike",
		Email: "mike@example.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	user.Email = "mike@anothercompany.com"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	userByEmail, err := us.ByEmail("mike@anothercompany.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(userByEmail)

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)
}
