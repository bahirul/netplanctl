package cmd

import (
	"fmt"
	"strconv"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/bahirul/netplanctl/internal/validation"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new network interface",
}

var createVlanCmd = &cobra.Command{
	Use:   "vlan",
	Short: "create a new VLAN interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("% command incomplete: create vlan <interface> <id>")
			return
		}

		parentIface := args[0]
		vlanId := args[1]
		ifaceName := parentIface + "." + vlanId

		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		if _, exists := config.Network.Vlans[ifaceName]; exists {
			fmt.Println("% VLAN interface", ifaceName, "already exists in the configuration.")
			return
		}

		if _, exists := config.Network.Ethernets[parentIface]; !exists {
			fmt.Println("% parent interface", parentIface, "not found in the configuration.")
			return
		}

		if err := validation.ValidateVlan(vlanId); err != nil {
			fmt.Println("% invalid VLAN ID:", err)
			return
		}

		vlanIdInt, err := strconv.Atoi(vlanId)
		if err != nil {
			fmt.Println("% invalid VLAN ID:", err)
			return
		}

		config.Network.Vlans[ifaceName] = netplan.NetplanVlan{
			Id:        vlanIdInt,
			Link:      parentIface,
			Addresses: []string{},
		}

		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

func init() {
	createCmd.AddCommand(createVlanCmd)
}
