package utils

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/pkg/errors"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func QueryAccountInfo(info keys.Info, cli rpcclient.Client) (uint64, uint64, error) {
	result, err := cli.ABCIQuery("/store/acc/key", auth.AddressStoreKey(info.GetAddress()))
	if err != nil {
		return 0, 0, err
	}

	if result.Response.Code == 0 {
		account := auth.BaseAccount{}
		err := GetTxCodec().UnmarshalBinaryBare(result.Response.Value, &account)
		if err != nil {
			return 0, 0, err
		}
		return account.AccountNumber, account.Sequence, nil
	}
	return 0, 0, errors.New("abci_query_error")
}
