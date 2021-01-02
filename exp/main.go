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
	// us.DestructiveReset()
	// user := models.User{
	// 	Name:  "Mike",
	// 	Email: "mike@example.com",
	// }
	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }
	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
