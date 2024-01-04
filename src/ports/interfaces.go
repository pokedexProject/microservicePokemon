package ports

import (
	model "github.com/pokedexProject/microservicePokemon/src/dominio"
)

// puerto de salida
type PokeRepository interface {
	CrearPokemon(input model.CrearPokemonInput) (*model.Pokemon, error)
	Pokemon(id string) (*model.Pokemon, error)
	ActualizarPokemon(id string, input *model.ActualizarPokemonInput) (*model.Pokemon, error)
	EliminarPokemon(id string) (*model.EliminacionPokemon, error)
	Pokemones() ([]*model.Pokemon, error)
}
