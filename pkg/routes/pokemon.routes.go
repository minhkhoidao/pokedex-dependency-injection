package routes

import (
	"poke-go/pkg/controllers"
	"poke-go/pkg/services"

	"github.com/gin-gonic/gin"
)

func RegisterPokemonRoutes(router *gin.Engine, service *services.PokemonService) {
	controller := controllers.NewPokemonController(service)

	pokemonGroups := router.Group("/pokemon")

	{
		pokemonGroups.POST("/create", controller.CreatePokemon)
	}
}
