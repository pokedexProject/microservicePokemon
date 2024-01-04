package service

import (
	"context"
	"fmt"

	model "github.com/pokedexProject/microservicePokemon/src/dominio"
	repository "github.com/pokedexProject/microservicePokemon/src/ports"
	pb "github.com/pokedexProject/microservicePokemon/src/proto"
)

// este servicio implementa la interfaz PokemonServiceServer
// que se genera a partir del archivo proto
type PokemonService struct {
	pb.UnimplementedPokemonServiceServer
	repo repository.PokeRepository
}

func NewPokemonService(repo repository.PokeRepository) *PokemonService {
	return &PokemonService{repo: repo}
}

func (s *PokemonService) CreatePokemon(ctx context.Context, req *pb.CreatePokemonRequest) (*pb.CreatePokemonResponse, error) {

	crearPokemonInput := model.CrearPokemonInput{
		Nombre:      req.GetNombre(),
		Tipo:        req.GetTipo(),
		Nivel:       req.GetNivel(),
		PuntosSalud: req.GetPuntosSalud(),
		Movimiento:  req.GetMovimiento(),
	}
	u, err := s.repo.CrearPokemon(crearPokemonInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Pokemon creado: %v", u)
	response := &pb.CreatePokemonResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Tipo:        u.Tipo,
		Nivel:       u.Nivel,
		PuntosSalud: u.PuntosSalud,
		Movimiento:  u.Movimiento,
	}
	fmt.Printf("Pokemon creado: %v", response)
	return response, nil
}

func (s *PokemonService) GetPokemon(ctx context.Context, req *pb.GetPokemonRequest) (*pb.GetPokemonResponse, error) {
	u, err := s.repo.Pokemon(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetPokemonResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Tipo:        u.Tipo,
		Nivel:       u.Nivel,
		PuntosSalud: u.PuntosSalud,
		Movimiento:  u.Movimiento,
	}
	return response, nil
}

func (s *PokemonService) ListPokemons(ctx context.Context, req *pb.ListPokemonsRequest) (*pb.ListPokemonsResponse, error) {
	pokemones, err := s.repo.Pokemones()
	if err != nil {
		return nil, err
	}
	var response []*pb.Pokemon
	for _, u := range pokemones {
		pokemon := &pb.Pokemon{
			Id:          u.ID,
			Nombre:      u.Nombre,
			Tipo:        u.Tipo,
			Nivel:       u.Nivel,
			PuntosSalud: u.PuntosSalud,
			Movimiento:  u.Movimiento,
		}
		response = append(response, pokemon)
	}

	return &pb.ListPokemonsResponse{Pokemons: response}, nil
}

func (s *PokemonService) UpdatePokemon(ctx context.Context, req *pb.UpdatePokemonRequest) (*pb.UpdatePokemonResponse, error) {
	nombre := req.GetNombre()
	tipo := req.GetTipo()
	nivel := req.GetNivel()
	puntosSalud := req.GetPuntosSalud()
	movimiento := req.GetMovimiento()
	fmt.Printf("Nombre: %v", nombre)
	actualizarPokemonInput := &model.ActualizarPokemonInput{
		Nombre:      &nombre,
		Tipo:        &tipo,
		Nivel:       &nivel,
		PuntosSalud: &puntosSalud,
		Movimiento:  &movimiento,
	}
	fmt.Printf("Pokemon actualizado input: %v", actualizarPokemonInput)
	u, err := s.repo.ActualizarPokemon(req.GetId(), actualizarPokemonInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Pokemon actualizado: %v", u)
	response := &pb.UpdatePokemonResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Tipo:        u.Tipo,
		Nivel:       u.Nivel,
		PuntosSalud: u.PuntosSalud,
		Movimiento:  u.Movimiento,
	}
	return response, nil
}

func (s *PokemonService) DeletePokemon(ctx context.Context, req *pb.DeletePokemonRequest) (*pb.DeletePokemonResponse, error) {
	respuesta, err := s.repo.EliminarPokemon(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeletePokemonResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}
