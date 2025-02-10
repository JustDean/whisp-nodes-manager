package node_api

import (
	"context"
	"log"
	"net"

	"github.com/JustDean/whisp-nodes-manager/pkg"
	"google.golang.org/grpc"
)

type NodeApi struct {
	listener   net.Listener
	manager    nodesManagerInterface
	grpcServer *grpc.Server
	UnimplementedNodeApiServer
}

func (server *NodeApi) Run(ctx context.Context) {
	log.Printf("Starting node-api server on %s", server.listener.Addr())
	go func() {
		server.grpcServer.Serve(server.listener)
	}()
	<-ctx.Done()
	log.Println("Stopping grpc server")
	server.grpcServer.GracefulStop()
}

func (server *NodeApi) Register(ctx context.Context, data *RegisterRequest) (*RegisterResponse, error) {
	nodeid, secret, err := server.manager.RegisterNode(data.GetInternalApiEndpoint(), data.GetPublicEndpoint())
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{
		NodeId: string(nodeid),
		Secret: secret,
	}, nil
}

func (server *NodeApi) Unregister(ctx context.Context, data *UnregisterRequest) (*UnregisterResponse, error) {
	server.manager.UnregisterNode(pkg.StringUUID(data.GetNodeId()))
	return &UnregisterResponse{}, nil
}

func SetServer(config pkg.GrpcServerConfig, nm nodesManagerInterface) (*NodeApi, error) {
	lis, err := net.Listen("tcp", config.Url())
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	server := &NodeApi{
		listener:   lis,
		manager:    nm,
		grpcServer: s,
	}
	RegisterNodeApiServer(s, server)
	return server, nil
}
