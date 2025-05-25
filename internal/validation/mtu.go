package validation

import (
	"fmt"
	"strconv"
)

func ValidateMtu(mtu string) error {
	if mtu == "" {
		return fmt.Errorf("MTU cannot be empty")
	}

	// Check if MTU is a valid number
	if _, err := strconv.Atoi(mtu); err != nil {
		return fmt.Errorf("invalid MTU value: %s", mtu)
	}

	// Check if MTU is within a reasonable range
	mtuValue, _ := strconv.Atoi(mtu)
	if mtuValue < 68 || mtuValue > 9000 {
		return fmt.Errorf("MTU value out of range (68-9000): %d", mtuValue)
	}

	return nil
}
