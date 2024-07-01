package database

import (
    "golang-login/models"
    "log"
)

func MigrateDB() {
    err := DB.AutoMigrate(&models.User{}, &models.Photo{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}
