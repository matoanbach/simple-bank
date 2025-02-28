package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/matoanbach/simple-bank/api"
	db "github.com/matoanbach/simple-bank/db/sqlc"
	"github.com/matoanbach/simple-bank/db/util"
	"github.com/matoanbach/simple-bank/gpai"
	"github.com/matoanbach/simple-bank/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

// evans --host localhost --port 9090 -r repl --package pb --service SimpleBank
func runGrpcServer(config util.Config, store db.Store) {
	server, err := gpai.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create the server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GPRCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
