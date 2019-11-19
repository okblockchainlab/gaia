package staking

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetStakingCmd() *cobra.Command {
	var stakingCmd = &cobra.Command{
		Use:   "staking",
		Short: "do something for validator or delegation",
		Long:  "do something for validator or delegation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("staking hahaha")
		},
	}
	return stakingCmd
}
