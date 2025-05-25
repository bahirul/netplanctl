package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Global variables
var version string
var netplanFile string
var temporaryNetplanFile string
var lastNetplanFile string

var rootCmd = &cobra.Command{
	Use:   "netplanctl",
	Short: "CLI to manage netplan configuration",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("% error loading .env file:", err)
		os.Exit(1)
	}

	netplanFileDefault := os.Getenv("NETPLAN_FILE")
	if netplanFileDefault == "" {
		netplanFileDefault = "/etc/netplan/50-cloud-init.yaml" // Default netplan file path
	}

	temporaryNetplanFile = os.Getenv("TEMPORARY_NETPLAN_FILE")
	if temporaryNetplanFile == "" {
		temporaryNetplanFile = "/etc/netplanctl/temp-netplan.yaml" // Default temporary netplan file path
	}
	lastNetplanFile = os.Getenv("LAST_NETPLAN_FILE")
	if lastNetplanFile == "" {
		lastNetplanFile = "/etc/netplanctl/last-netplan.yaml" // Default last netplan file path
	}

	if version == "" {
		version = "dev"
	}

	rootCmd.PersistentFlags().StringVarP(&netplanFile, "file", "f", netplanFileDefault, "path to the netplan configuration file")
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(shutdownCmd)
	rootCmd.AddCommand(startupCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(frrCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(createCmd)
}
