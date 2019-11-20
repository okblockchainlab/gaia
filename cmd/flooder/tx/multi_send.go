package tx

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/gaia/cmd/flooder/utils"
	"strings"
)

func NewMultiSendMsg(from keys.Info, tos []keys.Info, coin sdk.Coin) sdk.Msg {
	n := len(tos)
	inputCoinTotal := utils.CoinMul(coin, n)
	inputs := make([]bank.Input, 1)
	inputs[0] = bank.NewInput(from.GetAddress(), sdk.NewCoins(inputCoinTotal))
	outputs := make([]bank.Output, n)
	for i := 0; i < n; i++ {
		outputs[i] = bank.NewOutput(tos[i].GetAddress(), sdk.NewCoins(coin))
	}

	return bank.NewMsgMultiSend(inputs, outputs)
}

func MultiSend(from keys.Info, tos []keys.Info, coin sdk.Coin, nodeURI string, kb keys.Keybase) error {
	rpcClient := utils.GetRPCClient(nodeURI)
	defer rpcClient.Stop()

	msgs := NewMultiSendMsg(from, tos, coin)

	accountNumber, sequenceNum, err := QueryAccount(fromInfo, rpcClient)
	if err != nil {
		fmt.Println("query_account_error: ", err.Error())
	}
	//fmt.Printf("accountNumber: %d, sequenceNum: %d\n", accountNumber, sequenceNum)

	txBldr := GetTxBuilder(0, 0, kb)
	txBytes := BuildAndSign(fromInfo.GetName(), msgs, accountNumber, sequenceNum, txBldr)
	resp, err := rpcClient.BroadcastTxCommit(txBytes)
	if err != nil {
		//fmt.Println("BroadcastTxCommit_error: ", err)
		return err
	} else {
		//resp.CheckTx()
		if resp.DeliverTx.Code == 0 && strings.HasPrefix(resp.DeliverTx.Log, `[{"msg_index":0,"success":true,"`) {
			//fmt.Println("msgType:", msgs.Type(), resp.DeliverTx.Log)
			return nil
		} else {
			//fmt.Println("multi_send_error: ", resp)
			return errors.New(resp.DeliverTx.Log)
		}
	}
}
