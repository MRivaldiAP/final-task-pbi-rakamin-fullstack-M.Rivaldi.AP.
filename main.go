package main
import (
    "golang-login/databases"
    "golang-login/router"
)

func main() {
    database.ConnectDB()
    database.MigrateDB()
    r := router.SetupRouter()
    r.Run()
}