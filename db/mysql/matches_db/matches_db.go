package matches_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8",
		username, password, host,
	)

	log.Println(fmt.Sprintf("about to connect to %s\n", dataSource))
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

	_, err = Client.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v;", schema))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}

	_, err = Client.Exec(fmt.Sprintf("USE %v;", schema))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully using database..")
	}

	stmt, err := Client.Prepare("CREATE TABLE IF NOT EXISTS matchers(matchId VARCHAR(100) NOT NULL,seriesId VARCHAR(100) NOT NULL,teams VARCHAR(100) NOT NULL,details VARCHAR(200) NULL,url VARCHAR(200) NOT NULL,PRIMARY KEY(matchId));")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

	//defer Client.Close()

	/*
		To Do:
		load Match URLs from DB on start
	*/

}
