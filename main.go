package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Foo struct {
	ID          int `gorm:"primary_key"`
	Description string
}

var (
	DSN string = "root:root@tcp(localhost:3306)/db"
)

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// create first record
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Foo{Description: "foo"}).Error; err != nil {
			return err
		}

		log.Printf("foo created")

		return nil
	}); err != nil {
		log.Fatalf("failed to run transaction: %v", err)
	}

	// create second record (should rollback)
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Foo{Description: "bar"}).Error; err != nil {
			return err
		}

		return fmt.Errorf("should rollback")
	}); err != nil {
		log.Printf("2nd transaction returned error: %v", err)
	}

	// print records
	var foos []Foo
	if err := db.Find(&foos).Error; err != nil {
		log.Fatalf("failed to find foos: %v", err)
	}

	log.Printf("foos: %+v", foos)
}
