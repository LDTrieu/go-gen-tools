package main

import (
	"context"
	"fmt"
	"go-gen-tools/internal/repositories"
	"go-gen-tools/utils/metadata"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {

	// repo
	userRepo repositories.UserRepository
	 teamRepo repositories.TeamRepository

	// // service
	// userSrv services.UserService
	// authSrv services.AuthService
	// teamSrv services.TeamService

	// // deliveries
	// authpb pb.AuthServiceServer
	// userpb pb.UserServiceServer
	// teampb pb.TeamServiceServer

	// other
	TokenKey string
	mgoDB    *mongo.Database
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(); err != nil {
		return err
	}

	if err := srv.connectMongo(); err != nil {
		return err
	}

	if err := srv.loadRepositories(); err != nil {
		return err
	}

	if err := srv.loadServices(); err != nil {
		return err
	}

	if err := srv.loadDeliveries(); err != nil {
		return err
	}
	srv.startGRPCServer(ctx)
	return nil
}
func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) connectMongo() error {
	return nil

}
func (s *server) loadRedis() error {
	return nil
}
func (s *server) loadRepositories() error {
	// userRepo
	// teamRepo
	return nil
}

func (s *server) loadServices() error {
	// user
	// auth
	// team
	return nil
}

func (s *server) loadDeliveries() error {
	//s.authpb =
	//s.userpb =

	return nil

}
func (s *server) loadLogger() error {
	return nil
}

func (s *server) loadConfig() error {
	return nil
}

func (s *server) startGRPCServer(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(metadata.Authentication),
	)
	// if err := pb.RegisterAuthServiceHandlerServer(ctx, mux, s.authpb); err != nil {
	// 	return err
	// }

	// if err := pb.RegisterUserServiceHandlerServer(ctx, mux, s.userpb); err != nil {
	// 	return err
	// }

	// if err := pb.RegisterTeamServiceHandlerServer(ctx, mux, s.teampb); err != nil {
	// 	return err
	// }

	port := os.Getenv("PORT")
	log.Printf("server listen on port: %s\n", port)
	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
