package cmd

import (
	"fmt"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/spf13/cobra"
)

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "shutdown the network interfaces",
}

var ethernetShutdownCmd = &cobra.Command{
	Use:   "ethernet",
	Short: "shutdown an ethernet interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: shutdown ethernet <interface>")
			return
		}

		ifaceName := args[0]
		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		iface, ok := config.Network.Ethernets[ifaceName]
		if !ok {
			fmt.Println("% interface not found in the configuration:", ifaceName)
			return
		}

		// update the configuration
		iface.ActivationMode = "off"
		config.Network.Ethernets[ifaceName] = iface
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

var vlanShutdownCmd = &cobra.Command{
	Use:   "vlan",
	Short: "shutdown a VLAN interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: shutdown vlan <interface>")
			return
		}

		ifaceName := args[0]
		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		iface, ok := config.Network.Vlans[ifaceName]
		if !ok {
			fmt.Println("% interface not found in the configuration:", ifaceName)
			return
		}

		// update the configuration
		iface.ActivationMode = "off"
		config.Network.Vlans[ifaceName] = iface
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

func init() {
	shutdownCmd.AddCommand(ethernetShutdownCmd)
	shutdownCmd.AddCommand(vlanShutdownCmd)
}
