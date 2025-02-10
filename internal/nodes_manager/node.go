package nodes_manager

import (
	"context"
	"log"
	sync "sync"
	"time"

	"github.com/JustDean/whisp-nodes-manager/pkg"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// remote node's gRPC client (proxy to node)
type node struct {
	id                  pkg.StringUUID
	internalApiEndpoint string
	publicApiEndpoint   string
	closeCh             chan interface{}
	// nClients            int
	// hostedChats         []pkg.StringUUID
	nodeConn *grpc.ClientConn
	registry *nodesRegistry
}

func (n *node) hostChat(c ChatData) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := NewWhispNodeClient(n.nodeConn).HostChat(
		ctx,
		&HostChatRequest{
			ChatId:        string(c.Id),
			OwnerUsername: c.Owner,
		},
	)
	return err
}

func (n *node) dropChat(chatid pkg.StringUUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := NewWhispNodeClient(n.nodeConn).DropChat(
		ctx,
		&DropChatRequest{
			ChatId: string(chatid),
		},
	)
	return err
}

func (n *node) run(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
run:
	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			// TODO handle response
			_, err := NewWhispNodeClient(n.nodeConn).NodeInfo(ctx, &Blank{})
			if err != nil {
				n.registry.remove(n.id)
			}
			cancel()
		case <-ctx.Done():
			break run
		case <-n.closeCh:
			break run
		}
	}
	n.nodeConn.Close()
	log.Printf("Node %s is stoped", n.id)
}

func (n *node) stop() {
	n.closeCh <- struct{}{}
}

type nodesRegistry struct {
	registry     map[pkg.StringUUID]*node
	registerCh   chan *node
	unregisterCh chan pkg.StringUUID
	nextNodeidx  *struct {
		mu sync.Mutex
		v  int
	}
}

func (nr *nodesRegistry) run(ctx context.Context) {
	log.Println("Node registry is started")
	for {
		select {
		case newNode := <-nr.registerCh:
			nr.registry[newNode.id] = newNode
			go newNode.run(ctx)
			log.Printf("New node is started. Id %s ", newNode.id)
		case nodeid := <-nr.unregisterCh:
			node := nr.registry[nodeid]
			node.stop()
			delete(nr.registry, nodeid)
			log.Printf("Node is stopped. Id %s ", nodeid)
		case <-ctx.Done():
			log.Println("Nodes registry is stopped")
			return
		}
	}
}

func (nr *nodesRegistry) add(internalApiEndpoint, publicApiEndpoint string) (pkg.StringUUID, error) {
	conn, err := grpc.NewClient(internalApiEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// TODO handle response
	_, err = NewWhispNodeClient(conn).NodeInfo(ctx, &Blank{})
	if err != nil {
		return "", err
	}
	newNode := &node{
		id:                  pkg.StringUUID(uuid.NewString()),
		internalApiEndpoint: internalApiEndpoint,
		publicApiEndpoint:   publicApiEndpoint,
		closeCh:             make(chan interface{}),
		nodeConn:            conn,
		registry:            nr,
	}
	nr.registerCh <- newNode
	return newNode.id, nil
}

func (nr *nodesRegistry) remove(nodeid pkg.StringUUID) {
	_, ok := nr.registry[nodeid]
	if !ok {
		return
	}
	nr.unregisterCh <- nodeid
}

func (nr *nodesRegistry) chooseNode() *node {
	nr.nextNodeidx.mu.Lock()
	defer nr.nextNodeidx.mu.Unlock()
	registryLen := len(nr.registry)
	if nr.nextNodeidx.v > registryLen {
		nr.nextNodeidx.v = 0
	}
	i := 0
	for _, node := range nr.registry {
		if i == nr.nextNodeidx.v {
			return node
		}
		i++
	}
	nr.nextNodeidx.v++
	log.Println("No available node is found")
	return nil
}

func newNodesRegistry() *nodesRegistry {
	return &nodesRegistry{
		registry:     make(map[pkg.StringUUID]*node),
		registerCh:   make(chan *node),
		unregisterCh: make(chan pkg.StringUUID, 10),
		nextNodeidx: &struct {
			mu sync.Mutex
			v  int
		}{},
	}
}
