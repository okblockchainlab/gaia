package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCoinMul(t *testing.T) {
	coin := sdk.NewCoin("btc",sdk.NewInt(1024))
	retCoin:=CoinMul(coin,10)
	require.Equal(t,sdk.NewInt(10240),retCoin.Amount)
}
