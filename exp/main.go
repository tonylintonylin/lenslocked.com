package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "lenslocked_dev"
)

// gorm.Model is embeded, not inherited
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	// user := User{
	// 	Model: gorm.Model{
	// 		ID:        1,
	// 		CreatedAt: time.Now(),
	// 	},
	// }
	// fmt.Println(user.CreatedAt)
	
	//db.DropTableIfExists(&Ubers{})
	db.LogMode(true)
	db.AutoMigrate(&User{})
}
