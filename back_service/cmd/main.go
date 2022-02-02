package main

import (
	"fmt"
	"log"
	"net"

	backservice "github.com/ameena3/notch_project/back_service"
	"github.com/ameena3/notch_project/back_service/gen/product"
	dbClient "github.com/ameena3/notch_project/back_service/services/database_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Bootstrapping the service.
func main() {
	fmt.Println("Strating the service...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	ser := grpc.NewServer()
	client := dbClient.NewClient()
	backServer, err := backservice.NewServer(backservice.ServerParams{
		DbClient: client,
	})
	if err != nil {
		log.Fatalf("Failed to create new server: %v", err)
	}
	product.RegisterProductServiceServer(ser, backServer)
	reflection.Register(ser)
	log.Printf("server listening at %v", listener.Addr())
	if err := ser.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
