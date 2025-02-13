package nodes_manager

import (
	"context"
	"log"

	"github.com/JustDean/whisp-nodes-manager/pkg"
	"github.com/golang-jwt/jwt/v5"
)

// a facade for chats and nodes registries
type NodesManager struct {
	secret string
	nr     *nodesRegistry
	cr     *chatsRegistry
}

func (nm *NodesManager) Run(ctx context.Context) {
	log.Println("Starting Nodes Manager")
	go nm.nr.run(ctx)
	go nm.cr.run(ctx)
	<-ctx.Done()
	log.Println("Nodes manager is stopped")
}

func (nm NodesManager) RegisterNode(internalApiEndpoint, publicApiEndpoint string) (pkg.StringUUID, string, error) {
	nodeid, err := nm.nr.add(internalApiEndpoint, publicApiEndpoint)
	if err != nil {
		return "", "", err
	}
	return nodeid, nm.secret, nil
}

func (nm NodesManager) UnregisterNode(nodeid pkg.StringUUID) {
	nm.nr.remove(nodeid)
}

func (nm NodesManager) generateToken(chatid pkg.StringUUID, username string) string {
	encoder := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ChatId":   string(chatid),
		"Username": username,
	})
	token, _ := encoder.SignedString([]byte(nm.secret))
	return token
}

func (nm NodesManager) HostChat(chatName, chatPassword, ownerName string) (string, string, error) {
	node := nm.nr.chooseNode()
	newChat, err := nm.cr.hostChat(node, chatName, chatPassword, ownerName)
	if err != nil {
		return "", "", err
	}
	return nm.generateToken(newChat.Id, ownerName), node.publicApiEndpoint, nil
}

func (nm NodesManager) DropChat(chatid pkg.StringUUID, password string) {
	err := nm.cr.dropChat(chatid, password)
	if err != nil {
		log.Printf("Error droping chat %s: %v", chatid, err)
	}
}

func (nm NodesManager) ListChats() []ChatData {
	return nm.cr.listChats()
}

func (nm NodesManager) JoinChat(username, chatid, password string) (string, string, error) {
	chat, ok := nm.cr.getChat(pkg.StringUUID(chatid))
	if !ok {
		return "", "", CREDENTIALS_ERROR
	}
	if password != chat.password {
		return "", "", CREDENTIALS_ERROR
	}
	return nm.generateToken(chat.Id, username), chat.node.publicApiEndpoint, nil
}

func SetupNodesManager(config Config) *NodesManager {
	// TODO connect to db
	// perform recovery
	return &NodesManager{
		secret: config.Secret,
		nr:     newNodesRegistry(),
		cr:     newChatsRegistry(),
	}
}
