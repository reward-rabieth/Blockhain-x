package main

import (
	"context"
	node2 "github.com/reward-rabieth/blockchain/node"
	"github.com/reward-rabieth/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	node := node2.NewNode()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	log.Fatal(node.Start(":3000"))
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)
	version := &proto.Version{
		Version:    "blocker-0-1",
		Height:     1,
		ListenAddr: "4000",
	}
	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
