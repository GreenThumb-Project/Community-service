package main

import (
	"community-service/service"
	"community-service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "community-service/generated/community"
)

func main() {
	db, err := postgres.ConnectDb()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	community := service.CommunityServer{
		Community: postgres.NewCommunityRepo(db),
	}

	pb.RegisterComunityServiceServer(s, &community)

	log.Printf("server is running on %v...", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
