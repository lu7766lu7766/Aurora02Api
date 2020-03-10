package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)
import _ "github.com/jinzhu/gorm/dialects/mssql"

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_ip := os.Getenv("DB_IP")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pwd := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")

	var con string = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s;Initial Catalog=dbo;encrypt=disable;integrated security=SSPI", db_user, db_pwd, db_ip, db_port, db_name)
	// var con string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", db_ip, db_user, db_pwd, db_port, db_name)
	// fmt.Println(con)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "dbo." + defaultTableName 
	// }
	db, err := gorm.Open("mssql", con)
	if err != nil {
		// panic("failed to connect database")
		fmt.Println("error", err)
	}
	// db.SingularTable(true)
	return db
}
