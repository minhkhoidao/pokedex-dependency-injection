package main

import (
	"poke-go/pkg/database"
	"poke-go/pkg/repository"
	"poke-go/pkg/routes"
	"poke-go/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Initialize()
	PokemonRepository := repository.NewPokemonRepository()
	pokemonService := services.NewPokemonService(PokemonRepository)

	routes.RegisterPokemonRoutes(router, pokemonService)

	router.Run(":8080")
}
