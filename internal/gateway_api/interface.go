package gateway_api

import "github.com/JustDean/whisp-nodes-manager/pkg"

type nodesManagerInterface[Chat any] interface {
	HostChat(chatName, chatPassword, ownerName string) (token, connAddr string, err error)
	DropChat(pkg.StringUUID, string)
	JoinChat(chatId, chatPassword string) (token, connAddr string, err error)
	ListChats() []Chat
}
