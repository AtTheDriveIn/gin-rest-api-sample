package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Initialize(*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@/sample?charset=utf8&parseTime=True&loc=Local")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	// models.Migrate(db)

}
