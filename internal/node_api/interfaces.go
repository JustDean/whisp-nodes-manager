package node_api

import "github.com/JustDean/whisp-nodes-manager/pkg"

type nodesManagerInterface interface {
	RegisterNode(internalApiEndpoint, publicApiEndpoint string) (nodeid pkg.StringUUID, secret string, err error)
	UnregisterNode(pkg.StringUUID)
}
