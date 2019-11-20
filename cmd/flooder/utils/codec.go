package utils

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	txCodec = codec.New()
)

func init() {
	sdk.RegisterCodec(txCodec)
	txCodec.RegisterInterface((*crypto.PubKey)(nil), nil)
	txCodec.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
	auth.RegisterCodec(txCodec)
	bank.RegisterCodec(txCodec)
	staking.RegisterCodec(txCodec)
}

func GetTxCodec() *codec.Codec {
	return txCodec
}
