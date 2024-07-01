package controllers

import (
    "golang-login/databases"
    "golang-login/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// Struktur input untuk photo
type PhotoInput struct {
    Title    string `json:"title" binding:"required"`
    Caption  string `json:"caption" binding:"required"`
    PhotoUrl string `json:"photoUrl" binding:"required"`
}

// Fungsi untuk create photo
func CreatePhoto(c *gin.Context) {
    var input PhotoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    photo := models.Photo{Title: input.Title, Caption: input.Caption, PhotoUrl: input.PhotoUrl, UserID: userID.(uint)}
    database.DB.Create(&photo)

    c.JSON(http.StatusOK, gin.H{"message": "photo uploaded"})
}

// Fungsi untuk get photos
func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    database.DB.Preload("User").Find(&photos)
    c.JSON(http.StatusOK, photos)
}

// Fungsi untuk update photo
func UpdatePhoto(c *gin.Context) {
    photoID, err := strconv.Atoi(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid photo ID"})
        return
    }

    var photo models.Photo
    database.DB.First(&photo, photoID)
    if photo.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
        return
    }

    var input PhotoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID.(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to update this photo"})
        return
    }

    database.DB.Model(&photo).Updates(models.Photo{Title: input.Title, Caption: input.Caption, PhotoUrl: input.PhotoUrl})
    c.JSON(http.StatusOK, gin.H{"message": "photo updated"})
}

// Fungsi untuk delete photo
func DeletePhoto(c *gin.Context) {
    photoID, err := strconv.Atoi(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid photo ID"})
        return
    }

    var photo models.Photo
    database.DB.First(&photo, photoID)
    if photo.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID.(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to delete this photo"})
        return
    }

    database.DB.Delete(&photo)
    c.JSON(http.StatusOK, gin.H{"message": "photo deleted"})
}
