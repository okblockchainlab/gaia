package utils

import (
	"fmt"
	"github.com/cosmos/gaia/cmd/flooder/account"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	GaiaUrl = "tcp://localhost:26657"
)

func TestQueryAccountInfo(t *testing.T) {
	rpcClient := GetRPCClient(GaiaUrl)
	accountNum, SequenceNum, err := QueryAccountInfo(account.GetTuringInfo(), rpcClient)
	require.NoError(t, err)
	fmt.Println(accountNum,SequenceNum)
}
