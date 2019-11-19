package staking

import "github.com/spf13/cobra"

func GetStakingCmd() *cobra.Command {
	var stakingCmd = &cobra.Command{
		Use:   "staking",
		Short: "just 4 staking test",
		Long:  "just 4 staking test",
	}
	return stakingCmd
}
