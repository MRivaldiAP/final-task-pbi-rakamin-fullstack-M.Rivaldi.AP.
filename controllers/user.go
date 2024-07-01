package controllers

import (
    "golang-login/databases"
    "golang-login/models"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "os"
)

// Struktur input untuk registrasi
type RegisterInput struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// Fungsi untuk registrasi user
func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := models.User{Username: input.Username, Email: input.Email, Password: string(hashedPassword)}
    database.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

// Struktur input untuk login
type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// Fungsi untuk login user
func Login(c *gin.Context) {
    var input LoginInput
    var user models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Where("email = ?", input.Email).First(&user)
    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": user.ID,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })
    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Fungsi untuk update user
func UpdateUser(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var user models.User
    database.DB.First(&user, userID)
    if user.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    database.DB.Model(&user).Updates(models.User{Username: input.Username, Email: input.Email, Password: string(hashedPassword)})

    c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

// Fungsi untuk delete user
func DeleteUser(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var user models.User
    database.DB.First(&user, userID)
    if user.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
