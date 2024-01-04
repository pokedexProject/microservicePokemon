package dominio

type ActualizarPokemonInput struct {
	Nombre      *string `json:"nombre,omitempty"`
	Tipo        *string `json:"tipo,omitempty"`
	Nivel       *string `json:"nivel,omitempty"`
	PuntosSalud *string `json:"puntoSalud,omitempty"`
	Movimiento  *string `json:"movimiento,omitempty"`
}

type CrearPokemonInput struct {
	Nombre      string `json:"nombre"`
	Tipo        string `json:"tipo"`
	Nivel       string `json:"nivel"`
	PuntosSalud string `json:"puntoSalud"`
	Movimiento  string `json:"movimiento"`
}

type EliminacionPokemon struct {
	Mensaje string `json:"mensaje"`
}

type Pokemon struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Tipo        string `json:"tipo"`
	Nivel       string `json:"nivel"`
	PuntosSalud string `json:"puntoSalud"`
	Movimiento  string `json:"movimiento"`
}

func (Pokemon) IsEntity() {}
