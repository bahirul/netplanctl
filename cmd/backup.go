package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup the current netplan configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: backup <destination-file>")
			return
		}

		destPath := args[0]

		sourceFile, err := os.Open(netplanFile)
		if err != nil {
			fmt.Println("% error opening netplan file:", err)
			return
		}
		defer sourceFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			fmt.Println("% error creating backup file:", err)
			return
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			fmt.Println("% error copying netplan file to backup:", err)
			return
		}

		fmt.Println("backup created successfully at", destPath)
	},
}
