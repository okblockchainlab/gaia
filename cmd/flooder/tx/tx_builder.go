package tx

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	auth_utils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/gaia/cmd/flooder/account"
	"github.com/cosmos/gaia/cmd/flooder/utils"
)

const
(
	defaultGasAdjustment = float64(0)
	chainID              = "testchain"
	memo                 = ""

)

func GetTxBuilder(accountNum, sequenceNum uint64, kb keys.Keybase) auth.TxBuilder {
	txBuilder := auth.NewTxBuilder(
		auth_utils.GetTxEncoder(utils.GetTxCodec()),
		accountNum,
		sequenceNum,
		client.DefaultGasLimit,
		defaultGasAdjustment,
		false,
		chainID,
		memo,
		nil,
		nil,
	)
	txBuilder.WithKeybase(kb)
	return txBuilder
}

func BuildAndSign(fromName string, msg sdk.Msg, accountNum ,sequenceNum uint64, txBuilder auth.TxBuilder) []byte {
	txBuilder = txBuilder.WithAccountNumber(accountNum).WithSequence(sequenceNum)

	signMsg, err := txBuilder.BuildAndSign(fromName, account.DefaultEncryptPasswd, []sdk.Msg{msg})
	if err != nil {
		fmt.Println("buildAndSign error: ", err.Error())
	}
	return signMsg
}