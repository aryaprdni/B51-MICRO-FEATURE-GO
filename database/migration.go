package database

import (
	"B51-MICRO-FEATURE-GO/models"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Article{},
		&models.Paslon{},
		&models.Partai{},
		&models.User{},
		&models.Vote{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success")
}
