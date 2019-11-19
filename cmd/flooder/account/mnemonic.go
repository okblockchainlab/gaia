package account

import (
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
)

const entropySize = 256

var (
	mnemonicNum uint
)

func getMnemonicCmd() *cobra.Command {
	var mnemonicCmd = &cobra.Command{
		Use:   "mnemonic",
		Short: "generate mnemonics",
		Long:  "generate the account mnemonics",
		RunE:  generateMnemonics,
	}

	mnemonicCmd.Flags().UintVarP(&mnemonicNum, "number", "n", 1, "the number of mnemonics 2 generate")
	return mnemonicCmd
}

func generateMnemonics(*cobra.Command, []string) error {
	var i uint
	for ; i < mnemonicNum; i++ {
		seed, err := bip39.NewEntropy(entropySize)
		if err != nil {
			return err
		}

		mnemonic, err := bip39.NewMnemonic(seed)
		if err != nil {
			return err
		}
		fmt.Println(mnemonic)
	}
	return nil
}
