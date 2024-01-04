package dominio

import (
	"strconv"
)

// UsuarioGORM es el modelo de usuario para GORM de Usuario
type PokemonGORM struct {
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	Tipo        string `gorm:"type:varchar(255);not null"`
	Nivel       string `gorm:"type:varchar(255);not null"`
	PuntosSalud string `gorm:"type:varchar(255);not null"`
	Movimiento  string `gorm:"type:varchar(255);not null"`
}

// TableName especifica el nombre de la tabla para UsuarioGORM
func (PokemonGORM) TableName() string {
	return "pokemones"
}

func (pokemonGORM *PokemonGORM) ToGQL() (*Pokemon, error) {

	return &Pokemon{
		ID:          strconv.Itoa(int(pokemonGORM.ID)),
		Nombre:      pokemonGORM.Nombre,
		Tipo:        pokemonGORM.Tipo,
		Nivel:       pokemonGORM.Nivel,
		PuntosSalud: pokemonGORM.PuntosSalud,
		Movimiento:  pokemonGORM.Movimiento,
	}, nil
}
