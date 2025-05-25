package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit the changes to the netplan configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// ask the user to confirm
		fmt.Print("Are you sure you want to commit the changes? (yes/no): ")
		var response string
		fmt.Scanln(&response)
		if response != "yes" {
			fmt.Println("% commit aborted")
			return
		}

		tempConfig, err := netplan.RawConfig(temporaryNetplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		currentConfig, err := netplan.RawConfig(netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		dmp := diffmatchpatch.New()
		diff := dmp.DiffMain(currentConfig, tempConfig, false)

		if currentConfig == tempConfig || len(diff) == 0 {
			fmt.Println("% no changes found.")
			return
		}

		// load the temporary configuration file
		config, err := netplan.LoadConfig(temporaryNetplanFile)
		if err != nil {
			fmt.Println("%% error loading temporary configuration:", err)
			return
		}

		// backup the current netplan file copy it to a backup file
		sourceNetplanFile, err := os.Open(netplanFile)
		if err != nil {
			fmt.Println("% error getting netplan file:", err)
			return
		}
		defer sourceNetplanFile.Close()

		destNetplanFile, err := os.Create(lastNetplanFile)
		if err != nil {
			fmt.Println("% error creating backup file:", err)
			return
		}
		defer destNetplanFile.Close()

		_, err = io.Copy(destNetplanFile, sourceNetplanFile)
		if err != nil {
			fmt.Println("% error copying netplan file to backup:", err)
			return
		}

		// save the configuration to the netplan file
		if err := netplan.SaveConfig(config, netplanFile); err != nil {
			fmt.Println("% error saving changes:", err)
			return
		}

		// validate the netplan configuration
		validate := exec.Command("netplan", "generate", "--debug")
		if err := validate.Run(); err != nil {
			fmt.Println("% changes not committed, please fix the errors and try again")
			// restore the backup file
			if err := os.Rename(lastNetplanFile, netplanFile); err != nil {
				fmt.Println("% error restoring backup file:", err)
				return
			}
			fmt.Println("% rollback configuration from", lastNetplanFile)
			return
		}

		// apply the changes run cli netplan apply
		apply := exec.Command("netplan", "apply")
		if err := apply.Run(); err != nil {
			fmt.Println("% error applying changes:", err)
			return
		}

		// delete the temporary configuration file
		if err := os.Remove(temporaryNetplanFile); err != nil {
			fmt.Println("% error deleting temporary configuration file:", err)
			return
		}

		fmt.Println("changes committed successfully")
	},
}
