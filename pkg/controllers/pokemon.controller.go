package controllers

import (
	"net/http"
	"poke-go/pkg/model"
	"poke-go/pkg/services"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	service *services.PokemonService
}

func NewPokemonController(service *services.PokemonService) *PokemonController {
	return &PokemonController{
		service: service,
	}
}

func (bc *PokemonController) CreatePokemon(c *gin.Context) {
	var pkm model.Pokemon
	if err := c.ShouldBindJSON(&pkm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := bc.service.CreatePokemon(c, pkm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}
