package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CoinMul(origin sdk.Coin, multiple int) sdk.Coin {
	amount := origin.Amount.Mul(sdk.NewInt(int64(multiple)))
	return sdk.NewCoin(origin.Denom, amount)
}
