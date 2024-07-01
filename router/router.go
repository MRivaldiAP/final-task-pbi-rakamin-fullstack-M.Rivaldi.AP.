package router

import (
    "golang-login/controllers"
    "golang-login/middlewares"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // User routes
    router.POST("/users/register", controllers.Register)
    router.POST("/users/login", controllers.Login)
    router.PUT("/users/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
    router.DELETE("/users/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)

    // Photo routes
    router.POST("/photos", middlewares.AuthMiddleware(), controllers.CreatePhoto)
    router.GET("/photos", middlewares.AuthMiddleware(), controllers.GetPhotos)
    router.PUT("/photos/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)
    router.DELETE("/photos/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)

    return router
}
