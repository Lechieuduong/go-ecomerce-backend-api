package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Anh Duong Lap Trinh"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkOpenConns1(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:3307)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{

	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {

		}
	}

	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get db: %v", err)
	}

	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkOpenConns10(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:3307)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{

	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {

		}
	}

	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get db: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}