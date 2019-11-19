package main

import (
	"fmt"
	"github.com/cosmos/gaia/cmd/flooder/config"
	"github.com/cosmos/gaia/cmd/flooder/staking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	hosts         string
	broadcastMode string
)

var mainCmd = &cobra.Command{
	Use:   "flooder",
	Short: "A client tool for gaia-test",
	Long: `flooder is a very powerful cli tool for gaia-test. It supports JSON-file and Sub-command to stimulate transactions. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		return
	},
}

func init() {
	mainCmd.PersistentFlags().StringVarP(&hosts, config.FlagHosts, "u", config.DefaultHosts, "gaia's rcp port")
	mainCmd.PersistentFlags().StringVar(&broadcastMode, config.FlagBroadcastMode, config.DefaultBroadcastMode, "Transaction broadcasting mode (sync|async|block)")

	if err := viper.BindPFlag(config.FlagHosts, mainCmd.PersistentFlags().Lookup(config.FlagHosts)); err != nil {
		fmt.Println(err)
		return
	}

	if err := viper.BindPFlag(config.FlagBroadcastMode, mainCmd.PersistentFlags().Lookup(config.FlagBroadcastMode)); err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	mainCmd.AddCommand(staking.GetStakingCmd())
	Execute()
}

// Execute run the main cmd
func Execute() {
	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
