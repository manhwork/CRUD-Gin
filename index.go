package main

import "CRUD_Gin/router"

func main() {
	r := router.InitRouter()

	r.Run(":3000")
}
