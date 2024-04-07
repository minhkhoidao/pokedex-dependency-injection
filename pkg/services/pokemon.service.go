package services

import (
	"context"

	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/model"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/repository"
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

func (s *PokemonService) GetListPokemon(ctx context.Context) ([]model.Pokemon, error) {
	return s.repository.GetListPokemon(ctx)
}

func (s *PokemonService) GetPokemonById(ctx context.Context, id int) (*model.Pokemon, error) {
	return s.repository.GetPokemonById(ctx, id)
}

func (s *PokemonService) UpdatePokemon(ctx context.Context, pkm model.Pokemon) (*model.Pokemon, error) {
	return s.repository.UpdatePokemon(ctx, pkm)
}
