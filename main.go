package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID        int
	Plannings []Planning
}

type Planning struct {
	ID      int
	EventID int
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	// db.AutoMigrate(&Event{}, &Planning{})

	// db.Create(&Event{})
	// db.Create(&Event{})

	// db.Create(&Planning{EventID: 1})
	// db.Create(&Planning{EventID: 1})
	// db.Create(&Planning{EventID: 2})

	var events []Event

	// Doesn't work
	db.Preload("Plannings").Find(&events)
	log.Printf("%+v", events)

	// Works
	db.Preload("Plannings").Find(&events)
	log.Printf("%+v", events)
}
