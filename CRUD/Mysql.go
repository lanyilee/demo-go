package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type User struct {
	ID         int
	UserName   string
	Password   string
	UpdateTime time.Time
}

func (user *User) Insert(db *sql.DB) int64 {
	rs, err := db.Exec("Insert into t_User (UserName,Password,UpdateTime) values (?,?,?)", user.UserName, user.Password, user.UpdateTime)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	return id
}

var DB *sql.DB

func main() {
	DB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	user := &User{0, "张三", "12345", time.Now()}
	defer DB.Close()
	id := user.Insert(DB)
	fmt.Println(id)

}
