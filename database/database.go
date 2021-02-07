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
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbIP := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var con string = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s;Initial Catalog=dbo;encrypt=disable;integrated security=SSPI", dbUser, dbPwd, dbIP, dbPort, dbName)
	
	Eloquent, err = gorm.Open("mssql", con)

	if err != nil {
		// panic("failed to connect database")
		fmt.Println("db connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Println("database error %v", Eloquent.Error)
	}
	// fmt.Printf("%p, %T\n", Eloquent, Eloquent)
	// db.SingularTable(true)
	// return db
}
