package main

import (
	"tugas10/database"
	"tugas10/router"
)	

func main() {	

	database.StartDB()
	
	r:= router.StartApp()
	r.Run(":8080")
}