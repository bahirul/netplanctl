package validation

import (
	"fmt"
	"strconv"
)

func ValidateVlan(vlan string) error {
	if vlan == "" {
		return fmt.Errorf("VLAN ID cannot be empty")
	}

	vlanInt, err := strconv.Atoi(vlan)
	if err != nil {
		return fmt.Errorf("invalid VLAN ID: %s, must be a number", vlan)
	}

	// validate VLAN ID range
	if vlanInt < 1 || vlanInt > 4094 {
		return fmt.Errorf("invalid VLAN ID: %d, must be between 1 and 4094", vlanInt)
	}

	return nil
}
