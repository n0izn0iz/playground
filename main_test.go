package main

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type foo struct {
		gorm.Model
		BarCID string
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&foo{})
	if err != nil {
		t.Fatal("failed to migrate database")
	}

	// Create
	err = db.Create(&foo{BarCID: "some value"}).Error
	if err != nil {
		t.Fatal("failed to create model")
	}

	var m foo

	// Read with expected column name
	err = db.First(&m, "bar_cid = ?", "some value").Error // find with BarCID "some value"
	if err != nil {
		t.Fatal("can't fetch with expected column name")
	}

	// Read with actual column name
	err = db.First(&m, "bar_c_id = ?", "some value").Error // find with BarCID "some value"
	if err == nil {
		t.Fatal("can fetch with unexpected column name")
	}
}
