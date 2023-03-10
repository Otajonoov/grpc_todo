package main

import (
	pb "GRPC-TODO/genproto/store"
	"GRPC-TODO/server/postgres"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("Received : %v", in.GetName())
	store, err := postgres.CreateStore(&pb.Store{
		Id:          in.Id,
		Name:        in.Name,
		Discription: in.Discription,
		Addresses:   in.Addresses,
		IsOpen:      in.IsOpen,
	})

	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreServer) GetStore(ctx context.Context, in *pb.GetStoreRequest) (*pb.Store, error) {
	store, err := postgres.GetStore(in.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println(store)

	return store, nil
}

func (s *StoreServer) UpdateStore(ctx context.Context, in *pb.Store) (*emptypb.Empty, error) {
	err := postgres.UpdateStore(in)
	if err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
	return &emptypb.Empty{}, err
}

func (s *StoreServer) DeleteStore(ctx context.Context, in *pb.GetStoreRequest) (*emptypb.Empty, error) {
	err := postgres.DeleteStore(in.Id)
	if err != nil {
		log.Fatalf("Error to get store from server :%v", err)
	}

	return &emptypb.Empty{}, err
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed connection : %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStoreServiceServer(s, &StoreServer{})

	log.Printf("Server listening at : %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}

}
