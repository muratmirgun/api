package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID      int    `json:"id"`
	NAME    string `json:"name"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func main() {
	db := InitDB("../sqlite.db")
	AddUser(db, 2, "Murat")
	fmt.Println(GetUsers(db, 2))
}

func AddUser(db *sql.DB, ID int, NAME string) string {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into Users ( ID ,NAME ) values (?,?)")
	_, err := stmt.Exec(ID, NAME)
	checkError(err)
	tx.Commit()
	sum, _ := fmt.Println(ID)
	return string(rune(sum))
}

func GetUsers(db *sql.DB, id2 int) User {
	rows, err := db.Query("select * from Users")
	checkError(err)
	for rows.Next() {
		var tempUser User
		err =
			rows.Scan(&tempUser.ID, &tempUser.NAME)
		checkError(err)
		if tempUser.ID == id2 {
			return tempUser
		}
	}
	return User{}
}
