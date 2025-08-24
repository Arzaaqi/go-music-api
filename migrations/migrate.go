package migrations

import (
	"go-music-api/config"
	"go-music-api/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func RunMigrate() {
	err := config.DB.AutoMigrate(
		&models.Music{},
	)
	if err != nil {
		log.Fatal("Migration gagal:", err)
	}

	log.Println("Migration sukses 🚀")

}
