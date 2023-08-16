package node

import (
	"context"
	"fmt"
	"github.com/reward-rabieth/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	peer2 "google.golang.org/grpc/peer"
	"net"
	"sync"
)

type Node struct {
	version  string
	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blocker-0.1",
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	fmt.Printf("new peer connected (%s) - height (%d)\n", v.ListenAddr, v.Height)

	n.peers[c] = v

}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)

}

func (n *Node) Start(listenAddr string) error {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	fmt.Print("node running on port ", ":3000\n")
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)
	return grpcServer.Serve(ln)

}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}
	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}
	n.addPeer(c, v)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer2.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil

}
func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewNodeClient(c), nil
}
