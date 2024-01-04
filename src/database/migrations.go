package database

import (
	"log"

	model "github.com/pokedexProject/microservicePokemon/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.PokemonGORM{})

	log.Println("Migraciones completadas")
}
