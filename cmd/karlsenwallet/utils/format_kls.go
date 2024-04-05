package utils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/karlsen-network/karlsend/domain/consensus/utils/constants"
	"github.com/pkg/errors"
)

// FormatKls takes the amount of sompis as uint64, and returns amount of KLS with 8  decimal places
func FormatKls(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.SompiPerKarlsen)
	}
	return res
}

// KlsToSompi takes in a string representation of the Kls value to convert to Sompi
func KlsToSompi(amount string) (uint64, error) {
	err := validateKLSAmountFormat(amount)

	if err != nil {
		return 0, err
	}

	// after validation, amount can only be either an int OR
	// a float with an int component and decimal places
	parts := strings.Split(amount, ".")
	amountStr := ""

	if constants.SompiPerKarlsen%10 != 0 {
		return 0, errors.Errorf("Unable to convert to sompi when SompiPerKarlsen is not a multiple of 10")
	}

	decimalPlaces := int(math.Log10(constants.SompiPerKarlsen))
	decimalStr := ""

	if len(parts) == 2 {
		decimalStr = parts[1]
	}

	amountStr = fmt.Sprintf("%s%-*s", parts[0], decimalPlaces, decimalStr) // Padded with spaces at the end to fill for missing decimals: Sample "0.01234    "
	amountStr = strings.ReplaceAll(amountStr, " ", "0")                    // Make the spaces be 0s. Sample "0.012340000"

	convertedAmount, err := strconv.ParseUint(amountStr, 10, 64)

	return convertedAmount, err
}

func validateKLSAmountFormat(amount string) error {
	// Check whether it's an integer, or a float with max 8 digits
	match, err := regexp.MatchString("^([1-9]\\d{0,11}|0)(\\.\\d{0,8})?$", amount)

	if !match {
		return errors.Errorf("Invalid amount")
	}

	if err != nil {
		return err
	}

	return nil
}
