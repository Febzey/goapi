package main

import (
	"fmt"

	"github.com/febzey/goapi/pkg/configs"
	"github.com/febzey/goapi/pkg/routes"
	"github.com/febzey/goapi/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	routes.PublicRoutes(router)
	router.Use(mux.CORSMethodMiddleware(router))
	server := configs.ServerConf(router)
	utils.StartServer(server)
}
