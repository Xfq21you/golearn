package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func main() {
	// 连接数据库
	dsn := "root:@Xfq347@@tcp(127.0.0.1:3306)/test"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 检查连接
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully connected to the database!")

	// 创建表
	schema := `CREATE TABLE IF NOT EXISTS person (
        user_id INT AUTO_INCREMENT,
        username VARCHAR(260) NULL,
        sex VARCHAR(260) NULL,
        email VARCHAR(260) NULL,
        PRIMARY KEY (user_id)
    );`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Table person created or already exists.")

	// 插入数据
	person := Person{
		Username: "JohnDoe",
		Sex:      "Male",
		Email:    "john.doe@example.com",
	}

	insertQuery := `INSERT INTO person (username, sex, email) VALUES (:username, :sex, :email)`
	_, err = db.NamedExec(insertQuery, person)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Data inserted successfully!")

	// 查询数据
	people := []Person{}
	query := `SELECT user_id, username, sex, email FROM person`
	err = db.Select(&people, query)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Data retrieved from the database:")
	for _, p := range people {
		fmt.Printf("UserID: %d, Username: %s, Sex: %s, Email: %s\n", p.UserID, p.Username, p.Sex, p.Email)
	}
}
