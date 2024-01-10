package database

import (
	"B51-MICRO-FEATURE-GO/models"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Article{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success")
}
