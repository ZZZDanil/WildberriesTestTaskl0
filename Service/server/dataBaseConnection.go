package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DataBase struct {
	connection *sql.DB
}

func (DB *DataBase) ConnectToDataBase(dbo *DataBaseOptions) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbo.User, dbo.Password, dbo.DBName, dbo.SSLMode)
	fmt.Println(connStr)
	var err error
	DB.connection, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}

}
