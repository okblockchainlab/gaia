package account

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/cobra"
)

var (
	Kb         =keys.NewInMemory()
	turingInfo keys.Info
)



func GetAccountCmd() *cobra.Command {
	var accCmd = &cobra.Command{
		Use:   "account",
		Short: "about accounts",
		Long:  "about user accounts",
	}
	accCmd.AddCommand(getMnemonicCmd())
	accCmd.AddCommand(getMnemonicAndAddrToFile())
	return accCmd
}

func CreateTuring(kb keys.Keybase) (keys.Info, error) {
	return kb.CreateAccount(TuringName, TuringMnemonic, defaultBip39Passwd, DefaultEncryptPasswd, 0, 0)
}

func GetTuringInfo() keys.Info {
	return turingInfo
}
