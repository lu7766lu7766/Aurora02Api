package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var Eloquent *gorm.DB

func init() {
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
	
	Eloquent, err := gorm.Open("mssql", con)

	if err != nil {
		// panic("failed to connect database")
		fmt.Println("db connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Println("database error %v", Eloquent.Error)
	}
	fmt.Printf("%p, %T\n", Eloquent, Eloquent)
	// db.SingularTable(true)
	// return db
}
