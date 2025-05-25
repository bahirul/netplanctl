package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var frrCmd = &cobra.Command{
	Use:   "frr",
	Short: "interactive vtysh shell (frrouting)",
	Run: func(cmd *cobra.Command, args []string) {
		vtysh := exec.Command("vtysh")

		// Connect stdio to current terminal
		vtysh.Stdin = os.Stdin
		vtysh.Stdout = os.Stdout
		vtysh.Stderr = os.Stderr

		if err := vtysh.Run(); err != nil {
			fmt.Println("% failed to launch vtysh:", err)
		}
	},
}
