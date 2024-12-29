package main

import (
	"CRUD_Gin/database"
	"CRUD_Gin/router"
)

func main() {
	database.DBConn()

	r := router.InitRouter()

	r.Run(":3000")
}
