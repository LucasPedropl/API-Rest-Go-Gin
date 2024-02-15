package main

import (
	"gin/db"
	"gin/routes"
)

func main() {
	db.ConectaComBancoDeDados()
	routes.HandleRequest()
}
