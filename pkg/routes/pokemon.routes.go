package routes

import (
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/controllers"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/services"

	"github.com/gin-gonic/gin"
)

func RegisterPokemonRoutes(router *gin.Engine, service *services.PokemonService) {
	controller := controllers.NewPokemonController(service)

	pokemonGroups := router.Group("/pokemon")

	{
		pokemonGroups.GET("/", controller.GetListPokemon)
		pokemonGroups.GET("/:id", controller.GetPokemonById)
		pokemonGroups.POST("/create", controller.CreatePokemon)
		pokemonGroups.PUT("/update/:id", controller.UpdatePokemon)
	}
}
