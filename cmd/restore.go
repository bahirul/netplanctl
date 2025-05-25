package cmd

import (
	"fmt"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore the netplan configuration from a backup file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: restore <backup-file>")
			return
		}

		backupFile := args[0]

		config, err := netplan.LoadConfig(backupFile)
		if err != nil {
			fmt.Println("% error loading backup file:", err)
			return
		}

		if err := netplan.SaveConfig(config, netplanFile); err != nil {
			fmt.Println("% error restoring configuration:", err)
			return
		}

		fmt.Println("please run 'commit' to apply the changes.")
	},
}
