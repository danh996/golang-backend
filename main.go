package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/danh996/golang-backend/api"
	db "github.com/danh996/golang-backend/db/sqlc"
	"github.com/danh996/golang-backend/gapi"
	"github.com/danh996/golang-backend/pb"
	"github.com/danh996/golang-backend/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start GRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	srv, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can create server", err)
	}

	err = srv.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}

	fmt.Println("start server success")
}
