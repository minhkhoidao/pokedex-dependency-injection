package main

import (
	"log"

	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/database"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/repository"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/routes"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Initialize()
	PokemonRepository := repository.NewPokemonRepository()
	pokemonService := services.NewPokemonService(PokemonRepository)

	routes.RegisterPokemonRoutes(router, pokemonService)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
