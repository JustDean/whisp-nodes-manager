package gateway_api

import (
	"context"
	"log"
	"net"

	"github.com/JustDean/whisp-nodes-manager/internal/nodes_manager"
	"github.com/JustDean/whisp-nodes-manager/pkg"
	"google.golang.org/grpc"
)

type GatewayApi struct {
	listener   net.Listener
	manager    nodesManagerInterface[nodes_manager.ChatData]
	grpcServer *grpc.Server
	UnimplementedGatewayApiServer
}

func (server *GatewayApi) Run(ctx context.Context) {
	log.Printf("Starting GatewayApi server on %s", server.listener.Addr())
	go func() {
		server.grpcServer.Serve(server.listener)
	}()
	<-ctx.Done()
	log.Println("Stopping GatewayApi server")
	server.grpcServer.GracefulStop()
}

func (server *GatewayApi) HostChat(ctx context.Context, data *CreateChatRequest) (*ChatConnectionDataResponse, error) {
	token, hostAddr, err := server.manager.HostChat(data.GetName(), data.GetPassword(), data.GetOwnerUsername())
	if err != nil {
		return &ChatConnectionDataResponse{}, err
	}
	return &ChatConnectionDataResponse{
		Token:    token,
		Endpoint: hostAddr,
	}, nil
}

func (server *GatewayApi) DropChat(ctx context.Context, data *ChatCredentials) (*DropChatResponse, error) {
	server.manager.DropChat(pkg.StringUUID(data.GetId()), data.GetPassword())
	return &DropChatResponse{}, nil
}

func (server *GatewayApi) ListChats(ctx context.Context, data *ListChatsRequest) (*ListChatsResponse, error) {
	chats := server.manager.ListChats()
	var res []*ChatData
	for _, chat := range chats {
		res = append(res, &ChatData{
			Id:            string(chat.Id),
			Name:          chat.Name,
			OwnerUsername: chat.Owner,
		})
	}
	return &ListChatsResponse{Chats: res}, nil
}

func (server *GatewayApi) JoinChat(ctx context.Context, data *ChatCredentials) (*ChatConnectionDataResponse, error) {
	token, connAddr, err := server.manager.JoinChat(data.GetId(), data.GetPassword())
	if err != nil {
		return &ChatConnectionDataResponse{}, err
	}
	return &ChatConnectionDataResponse{Token: token, Endpoint: connAddr}, nil
}

func SetServer(config pkg.GrpcServerConfig, nm nodesManagerInterface[nodes_manager.ChatData]) (*GatewayApi, error) {
	lis, err := net.Listen("tcp", config.Url())
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	server := &GatewayApi{
		listener:   lis,
		manager:    nm,
		grpcServer: s,
	}
	RegisterGatewayApiServer(s, server)
	return server, nil
}
