package account

import "github.com/spf13/cobra"

func GetAccountCmd() *cobra.Command {
	var accCmd = &cobra.Command{
		Use:   "account",
		Short: "about accounts",
		Long:  "about user accounts",
	}
	accCmd.AddCommand(getMnemonicCmd())
	return accCmd
}
