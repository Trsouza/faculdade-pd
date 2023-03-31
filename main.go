package main

import (
    "fmt"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/server"
	_ "github.com/trsouza/faculdade-pd/docs"
)

// @title API Faculdade PD
// @version 1.0
// @description Documentação da API construída em Golang
// @contact.name Tathiane Souza
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /api/v1
func main() {

    fmt.Println("Carregando...")
	database.StartDB()
	s := server.NewServer()

	s.Run()
}