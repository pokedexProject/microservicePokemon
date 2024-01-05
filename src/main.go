package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	repository "github.com/pokedexProject/microservicePokemon/src/adapters"
	service "github.com/pokedexProject/microservicePokemon/src/aplication"
	"github.com/pokedexProject/microservicePokemon/src/database"
	pb "github.com/pokedexProject/microservicePokemon/src/proto"
	"google.golang.org/grpc"
)

func main() {

	db := database.Connect()
	database.EjecutarMigraciones(db.GetConn())
	pokemonRepository := repository.NewPokeRepository(db)
	pokemonService := service.NewPokemonService(pokemonRepository)
	// Configura el servidor gRPC
	//este servidor está escuchando en el puerto 50052
	//y se encarga de registrar el servicio de entrenadores
	grpcServe := grpc.NewServer()
	// Registra el servicio de entrenadores en el servidor gRPC
	pb.RegisterPokemonServiceServer(grpcServe, pokemonService)

	// Define el puerto en el que se ejecutará el servidor gRPC
	port := "50052"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server listening on port %s...\n", port)

	// Inicia el servidor gRPC en segundo plano
	go func() {
		if err := grpcServe.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Espera una señal para detener el servidor gRPC
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	fmt.Println("Shutting down the server...")

	// Detén el servidor gRPC de manera segura
	grpcServe.GracefulStop()
}
