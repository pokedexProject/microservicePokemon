package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/pokedexProject/microservicePokemon/src/database"
	model "github.com/pokedexProject/microservicePokemon/src/dominio"
	"github.com/pokedexProject/microservicePokemon/src/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type pokeRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewPokeRepository(db *database.DB) ports.PokeRepository {
	return &pokeRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func ToJSON(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// Retrieve obtiene un pokemon por su correo y contraseña.

// ObtenerTrabajo obtiene un trabajo por su ID.
func (ur *pokeRepository) Pokemon(id string) (*model.Pokemon, error) {
	if id == "" {
		return nil, errors.New("El ID de pokemon es requerido")
	}

	var pokemonGORM model.PokemonGORM
	//result := ur.db.GetConn().First(&pokemonGORM, id)
	result := ur.db.GetConn().First(&pokemonGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el pokemon con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return pokemonGORM.ToGQL()
}

// Pokemones obtiene todos los pokemons de la base de datos.
func (ur *pokeRepository) Pokemones() ([]*model.Pokemon, error) {
	var pokemonsGORM []model.PokemonGORM
	result := ur.db.GetConn().Find(&pokemonsGORM)

	if result.Error != nil {
		log.Printf("Error al obtener los pokemons: %v", result.Error)
		return nil, result.Error
	}

	var pokemons []*model.Pokemon
	for _, pokemonGORM := range pokemonsGORM {
		pokemon, _ := pokemonGORM.ToGQL()
		pokemons = append(pokemons, pokemon)
	}

	// pokemonsJSON, err := json.Marshal(pokemons)
	// if err != nil {
	// 	log.Printf("Error al convertir pokemons a JSON: %v", err)
	// 	return "[]", err
	// }
	// return ToJSON(pokemons)
	return pokemons, nil
}
func (ur *pokeRepository) CrearPokemon(input model.CrearPokemonInput) (*model.Pokemon, error) {
	pokemonGORM :=
		&model.PokemonGORM{
			Nombre:      input.Nombre,
			Tipo:        input.Tipo,
			Nivel:       input.Nivel,
			PuntosSalud: input.PuntosSalud,
			Movimiento:  input.Movimiento,
		}
	result := ur.db.GetConn().Create(&pokemonGORM)
	if result.Error != nil {
		log.Printf("Error al crear el pokemon: %v", result.Error)
		return nil, result.Error
	}

	response, err := pokemonGORM.ToGQL()
	return response, err
}

func (ur *pokeRepository) ActualizarPokemon(id string, input *model.ActualizarPokemonInput) (*model.Pokemon, error) {
	var pokemonGORM model.PokemonGORM
	if id == "" {
		return nil, errors.New("El ID de pokemon es requerido")
	}

	result := ur.db.GetConn().First(&pokemonGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Pokemon con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Solo actualiza los campos proporcionados
	if input.Nombre != nil {
		pokemonGORM.Nombre = *input.Nombre
	}
	if input.Tipo != nil {
		pokemonGORM.Tipo = *input.Tipo
	}
	if input.Nivel != nil {
		pokemonGORM.Nivel = *input.Nivel
	}
	if input.PuntosSalud != nil {
		pokemonGORM.PuntosSalud = *input.PuntosSalud
	}
	if input.Movimiento != nil {
		pokemonGORM.Movimiento = *input.Movimiento
	}
	result = ur.db.GetConn().Save(&pokemonGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Pokemon actualizado: %v", pokemonGORM)
	return pokemonGORM.ToGQL()
}

// EliminarPokemon elimina un pokemon de la base de datos por su ID.
func (ur *pokeRepository) EliminarPokemon(id string) (*model.EliminacionPokemon, error) {
	// Intenta buscar el pokemon por su ID
	var pokemonGORM model.PokemonGORM
	result := ur.db.GetConn().First(&pokemonGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El pokemon no se encontró en la base de datos
			response := &model.EliminacionPokemon{
				Mensaje: "El pokemon no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el pokemon con ID %s: %v", id, result.Error)
		response := &model.EliminacionPokemon{
			Mensaje: "Error al buscar el pokemon",
		}
		return response, result.Error
	}

	// Elimina el pokemon de la base de datos
	result = ur.db.GetConn().Delete(&pokemonGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el pokemon con ID %s: %v", id, result.Error)
		response := &model.EliminacionPokemon{
			Mensaje: "Error al eliminar el pokemon",
		}
		return response, result.Error
	}

	// Éxito al eliminar el pokemon
	response := &model.EliminacionPokemon{
		Mensaje: "Pokemon eliminado con éxito",
	}
	return response, result.Error

}
