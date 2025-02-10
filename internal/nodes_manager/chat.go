package nodes_manager

import (
	"context"

	"github.com/JustDean/whisp-nodes-manager/pkg"
	"github.com/google/uuid"
)

type ChatData struct {
	Id       pkg.StringUUID
	Name     string
	Owner    string
	password string
	node     *node
}

type chatsRegistry struct {
	registry     map[pkg.StringUUID]ChatData
	registerCh   chan ChatData
	unregisterCh chan pkg.StringUUID
}

func (cr *chatsRegistry) run(ctx context.Context) {
	for {
		select {
		case newChat := <-cr.registerCh:
			cr.registry[newChat.Id] = newChat
		case chatid := <-cr.unregisterCh:
			delete(cr.registry, chatid)
		case <-ctx.Done():
			return
		}
	}
}

func (cr chatsRegistry) listChats() []ChatData {
	var chats []ChatData
	for _, chat := range cr.registry {
		chats = append(chats, chat)
	}
	return chats
}

func (cr *chatsRegistry) hostChat(node *node, chatName, chatPassword, ownerName string) (ChatData, error) {
	newChat := ChatData{
		Id:       pkg.StringUUID(uuid.NewString()),
		Name:     chatName,
		Owner:    ownerName,
		password: chatPassword,
		node:     node,
	}
	err := node.hostChat(newChat)
	if err != nil {
		return ChatData{}, err
	}
	cr.registerCh <- newChat
	return newChat, nil
}

func (cr *chatsRegistry) dropChat(chatid pkg.StringUUID, password string) error {
	chat, ok := cr.registry[chatid]
	if !ok {
		return nil
	}
	if chat.password != password {
		return CREDENTIALS_ERROR
	}
	err := chat.node.dropChat(chatid)
	if err != nil {
		return err
	}
	cr.unregisterCh <- chatid
	return nil
}

func (cr chatsRegistry) getChat(chatid pkg.StringUUID) (ChatData, bool) {
	chat, ok := cr.registry[chatid]
	return chat, ok
}

func newChatsRegistry() *chatsRegistry {
	return &chatsRegistry{
		registry:     make(map[pkg.StringUUID]ChatData),
		registerCh:   make(chan ChatData),
		unregisterCh: make(chan pkg.StringUUID, 10),
	}
}
