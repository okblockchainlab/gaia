package tx

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gaia/cmd/flooder/account"
	"testing"
)

var (
	kb keys.Keybase
)

func init() {
	kb = keys.NewInMemory()
}

func TestBuildAndSign(t *testing.T) {
	turing, err := account.CreateTuring(kb)
	if err != nil {
		fmt.Println(err)
	}

	txBuilder := GetTxBuilder(0, 0, kb)

	tos := account.GenerateAccount(kb, 2)
	coin := sdk.NewCoin(DefaultDenom, sdk.NewInt(10))
	msg := NewMultiSendMsg(turing, tos, coin)

	fmt.Println(msg)

	bz := BuildAndSign(turing.GetName(), msg, 0, 0, txBuilder)
	fmt.Println(bz)
}
