package repository

import (
	"context"
	"errors"

	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/database"
	"github.com/minhkhoidao/pokedex-dependency-injection/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPokemonRepository interface {
	CreatePokemon(ctx context.Context, pokemon model.Pokemon) (*model.Pokemon, error)
	GetListPokemon(ctx context.Context) ([]model.Pokemon, error)
	GetPokemonById(ctx context.Context, id int) (*model.Pokemon, error)
	UpdatePokemon(ctx context.Context, book model.Pokemon) (*model.Pokemon, error)
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

func (r *PokemonRepository) GetListPokemon(ctx context.Context) ([]model.Pokemon, error) {
	var pokemons []model.Pokemon

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pokemon model.Pokemon
		cursor.Decode(&pokemon)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (r *PokemonRepository) GetPokemonById(ctx context.Context, id int) (*model.Pokemon, error) {
	var pokemon model.Pokemon

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&pokemon)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (r *PokemonRepository) UpdatePokemon(ctx context.Context, pkm model.Pokemon) (*model.Pokemon, error) {
	filter := bson.M{"_id": pkm.ID}
	update := bson.M{"$set": pkm}

	result, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	// Perform the update operation using the filter and update document
	if result.MatchedCount == 0 {
		return nil, errors.New("no book found with given ID")
	}

	// Check if the document was actually updated
	return &pkm, nil
}
