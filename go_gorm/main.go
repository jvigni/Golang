package main

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TestValue string
}

/* Custom table name:
// Override the table name
func (User) TableName() string {
	return "custom_users" // Specify the table name here
}
*/

func main() {
	host := "localhost"
	user := "juan"
	dbName := "postgres"
	password := "123"
	port := "5432"
	timezone := "UTC"
	dsn := "host=" + host + // Data Source Name
		" user=" + user +
		" dbname=" + dbName +
		" password=" + password +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=" + timezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db.Migrator().CreateTable(&User{})
	//db.AutoMigrate(&User{}) // Automatically migrate the schema, to keep it updated [refresh if it changed]
	db.Create(&User{TestValue: "test value3"})
}
