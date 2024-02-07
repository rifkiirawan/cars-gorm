package main

import (
	"cars-gorm/database"
	"cars-gorm/routers"
)

//UNTUK GIN

func main() {
	var PORT = ":8080"
	database.StartDB()

	routers.StartServer().Run(PORT)

}
