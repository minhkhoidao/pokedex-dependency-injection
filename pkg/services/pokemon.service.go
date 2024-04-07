package services

import (
	"context"
	"poke-go/pkg/model"
	"poke-go/pkg/repository"
)

type PokemonService struct {
	repository repository.IPokemonRepository
}

func NewPokemonService(repo repository.IPokemonRepository) *PokemonService {
	return &PokemonService{
		repository: repo,
	}
}

func (s *PokemonService) CreatePokemon(ctx context.Context, pkm model.Pokemon) (*model.Pokemon, error) {
	return s.repository.CreatePokemon(ctx, pkm)
}
