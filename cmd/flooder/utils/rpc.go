package utils

import(
	rpcclient "github.com/tendermint/tendermint/rpc/client"

)

func GetRPCClient(nodeURI string) (client rpcclient.Client) {
	return rpcclient.NewHTTP(nodeURI, "/websocket")
}