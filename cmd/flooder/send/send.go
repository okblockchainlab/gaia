package send

import "github.com/spf13/cobra"

func GetSendCmd() *cobra.Command {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "transfer 2 an address",
		Long:  "transfer 2 an address",
		RunE:  goSend,
	}

	return sendCmd
}

func goSend(*cobra.Command, []string) error {
	return nil
}
