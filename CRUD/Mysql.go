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

func (user *User) Update(db *sql.DB) {
	rs, err := db.Exec("update t_User SET Password = ?,UpdateTime=? WHERE UserName=?", user.Password, time.Now(), user.UserName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rs)
}

func Search(db *sql.DB) User {
	user := User{}
	err := db.QueryRow("select UserName,Password from t_User where  id = ?", 1).Scan(&user.UserName, &user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return user
}
func SearchList(db *sql.DB) (users []User, err error) {
	rows, err := db.Query("select UserName,Password from t_User ")
	for rows.Next() {
		user := User{}
		er := rows.Scan(&user.UserName, &user.Password)
		if er != nil {
			log.Fatalln(er)
		}
		users = append(users, user)
	}
	return
}

var DB *sql.DB

func main() {
	DB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	//user := &User{0, "张三", "12345", time.Now()}
	defer DB.Close()
	user := Search(DB)
	fmt.Println(user.Password)
	fmt.Println(user.UserName)
	users, err := SearchList(DB)
	for _, u := range users {
		fmt.Println(u.UserName + u.Password)
	}
	//id := user.Insert(DB)
	//fmt.Println(id)
	//user.Password = "abcde"
	//user.Update(DB)
}
