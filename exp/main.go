package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	Color string
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

	name, email, color := getInfo()
	u := User{
		Name:  name,
		Email: email,
		Color: color,
	}
	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func getInfo() (name, email, color string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("what is your name?")
	name, _ = reader.ReadString('\n')

	fmt.Println("what is your email?")
	email, _ = reader.ReadString('\n')

	fmt.Println("what is your color?")
	color, _ = reader.ReadString('\n')

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)

	return name, email, color
}
