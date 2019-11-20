package tx

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gaia/cmd/flooder/account"
	"testing"
)

const (
	DefaultDenom = "stake"
	GaiaUrl      = "tcp://localhost:26657"
)

func TestMultiSend(t *testing.T) {

	kb:=keys.NewInMemory()
	from ,err:= account.CreateTuring(kb)
	if err!=nil{
		fmt.Println(err)
		return
	}
	tos := account.GenerateAccount(kb, 10)
	coin := sdk.NewCoin(DefaultDenom, sdk.NewInt(10))
	/////////////////////////// notice~~~~~~~~~~~~//////////////////////
 	 // what the motherf*cker : buildAndSign error:  Key Turing not found'
	if err := MultiSend(from, tos, coin, GaiaUrl, kb); err != nil {
		fmt.Println(err)
	}
	/////////////////////////// notice~~~~~~~~~~~~//////////////////////

	//require.NoError(t, err)
}
