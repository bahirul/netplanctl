package cmd

import (
	"fmt"
	"os"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/bahirul/netplanctl/internal/validation"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set the configuration for a network interface",
}

var setEthernetCmd = &cobra.Command{
	Use:   "ethernet",
	Short: "set the configuration for a network interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: set ethernet <interface>")
			return
		}

		ipFlag := cmd.Flag("ip").Changed
		mtuFlag := cmd.Flag("mtu").Changed
		stateFlag := cmd.Flag("state").Changed

		if !ipFlag && !mtuFlag && !stateFlag {
			fmt.Println("% please provide at least one flag: --ip, --mtu, --state")
			cmd.Usage()
			return
		}

		ifaceName := args[0]
		ipAddr := cmd.Flag("ip").Value.String()
		mtu := cmd.Flag("mtu").Value.String()

		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)

		if err != nil {
			fmt.Println("% error:", err)
			os.Exit(1)
		}

		if _, ok := config.Network.Ethernets[ifaceName]; !ok {
			fmt.Println("% interface not found in the configuration:", ifaceName)
			return
		}

		iface := config.Network.Ethernets[ifaceName]

		if mtu != "" {
			// validate the MTU value
			if err := validation.ValidateMtu(mtu); err != nil {
				fmt.Println("% invalid MTU value:", err)
				return
			}
			iface.MTU = mtu
		}

		if ipAddr != "" {
			// validate the IP address format
			if err := validation.ValidateIP(ipAddr); err != nil {
				fmt.Println("% invalid IP address format:", err)
				return
			}

			iface.Addresses = append(iface.Addresses, ipAddr)

			// remove duplicate IP addresses
			ipMap := make(map[string]bool)
			for _, ip := range iface.Addresses {
				ipMap[ip] = true
			}
			iface.Addresses = []string{}
			for ip := range ipMap {
				iface.Addresses = append(iface.Addresses, ip)
			}

		}

		// update the configuration
		config.Network.Ethernets[ifaceName] = iface
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

var setVlanCmd = &cobra.Command{
	Use:   "vlan",
	Short: "set the configuration for a VLAN interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: set vlan <interface>")
			return
		}

		ipFlag := cmd.Flag("ip").Changed

		if !ipFlag {
			fmt.Println("% please provide at least one flag: --ip")
			cmd.Usage()
			return
		}

		ifaceName := args[0]
		ipAddr := cmd.Flag("ip").Value.String()

		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			os.Exit(1)
		}

		if _, ok := config.Network.Vlans[ifaceName]; !ok {
			fmt.Println("% interface not found in the configuration:", ifaceName)
			return
		}

		vlan := config.Network.Vlans[ifaceName]

		if ipAddr != "" {
			if err := validation.ValidateIP(ipAddr); err != nil {
				fmt.Println("% invalid IP address format:", err)
				return
			}
			vlan.Addresses = append(vlan.Addresses, ipAddr)

			ipMap := make(map[string]bool)
			for _, ip := range vlan.Addresses {
				ipMap[ip] = true
			}
			vlan.Addresses = []string{}
			for ip := range ipMap {
				vlan.Addresses = append(vlan.Addresses, ip)
			}
		}

		config.Network.Vlans[ifaceName] = vlan
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

func init() {
	setCmd.AddCommand(setEthernetCmd)
	setEthernetCmd.Flags().StringP("ip", "", "", "ip address to set")
	setEthernetCmd.Flags().StringP("mtu", "", "", "mtu to set")

	setCmd.AddCommand(setVlanCmd)
	setVlanCmd.Flags().StringP("ip", "", "", "ip address to set")
}
