package repository

import (
	"context"
	"poke-go/pkg/database"
	"poke-go/pkg/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type IPokemonRepository interface {
	CreatePokemon(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
	// GetListPokemon(ctx context.Context) ([]model.Pokemon, error)
	// GetPokemonById(ctx context.Context, id int) (*model.Pokemon, error)
}

type PokemonRepository struct {
	collection *mongo.Collection
}

func NewPokemonRepository() *PokemonRepository {
	return &PokemonRepository{
		collection: database.DB.Collection("pokemon"),
	}
}

func (r *PokemonRepository) CreatePokemon(ctx context.Context, pkm model.Pokemon) (*model.Pokemon, error) {
	_, err := r.collection.InsertOne(ctx, pkm)
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}
