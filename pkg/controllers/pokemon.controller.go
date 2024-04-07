package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/model"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/services"

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

func (bc *PokemonController) GetListPokemon(c *gin.Context) {
	result, err := bc.service.GetListPokemon(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (bc *PokemonController) GetPokemonById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return

	}
	log.Println(idInt)
	result, err := bc.service.GetPokemonById(c, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (bc *PokemonController) UpdatePokemon(c *gin.Context) {
	var pkm model.Pokemon
	if err := c.ShouldBindJSON(&pkm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := bc.service.UpdatePokemon(c, pkm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
