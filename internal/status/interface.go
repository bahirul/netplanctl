package status

import (
	"fmt"
	"os"
	"strings"
)

// GetInterfaceStatus reads operstate and carrier from sysfs
func GetInterfaceStatus(ifaceName string) (operState string, carrier string, err error) {
	read := func(path string) (string, error) {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(data)), nil
	}

	operState, err = read(fmt.Sprintf("/sys/class/net/%s/operstate", ifaceName))
	if err != nil {
		return "", "", fmt.Errorf("failed to read operstate: %w", err)
	}

	carrier, err = read(fmt.Sprintf("/sys/class/net/%s/carrier", ifaceName))
	if err != nil {
		return "", "", fmt.Errorf("failed to read carrier: %w", err)
	}

	return operState, carrier, nil
}

// GetInterfaceMTU reads the MTU from sysfs
func GetInterfaceMTU(ifaceName string) (mtu string, err error) {
	read := func(path string) (string, error) {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(data)), nil
	}

	mtu, err = read(fmt.Sprintf("/sys/class/net/%s/mtu", ifaceName))
	if err != nil {
		return "", fmt.Errorf("failed to read MTU: %w", err)
	}
	return mtu, nil
}
