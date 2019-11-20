package account

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/cobra"
)

var (
	kb         = keys.NewInMemory()
	turingInfo keys.Info
)

func init() {
	turingInfo, _ = createTuring(kb)
}

func GetAccountCmd() *cobra.Command {
	var accCmd = &cobra.Command{
		Use:   "account",
		Short: "about accounts",
		Long:  "about user accounts",
	}
	accCmd.AddCommand(getMnemonicCmd())
	return accCmd
}

func createTuring(kb keys.Keybase) (keys.Info, error) {
	return kb.CreateAccount(TuringName, TuringMnemonic, defaultBip39Passwd, defaultEncryptPasswd, 0, 0)
}

func GetTuringInfo() keys.Info {
	return turingInfo
}
