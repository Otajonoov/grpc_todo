package main

import (
	"context"
	"log"
	"time"

	pb "GRPC-TODO/genproto/store"

	"google.golang.org/grpc"
)

type Store struct {
	ID          int64
	Name        string
	Discription string
	Addresses   []string
	IsOpen      bool
}

const (
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewStoreServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// store, err := c.CreateStore(ctx, &pb.Store{
	// 	Name:        "Quvonchbek",
	// 	Discription: "Go",
	// 	IsOpen:      true,
	// 	Addresses: []string{
	// 		"Xon",
	// 		"Shovot",
	// 	},
	// })
	// fmt.Println(store)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	
	// _, err = c.UpdateStore(ctx, &pb.Store{})
	// if err != nil {
	// 	log.Fatalf("Error to update : %v", err)
	// }

	_, err = c.DeleteStore(ctx, &pb.GetStoreRequest{Id: 2})
	if err != nil {
		log.Fatalf("Error to delete from client :%v", err)
	}

}
