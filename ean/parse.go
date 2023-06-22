package ean

import (
	"eancheck/gs1"
	"fmt"
	"strconv"
	"strings"
)

// ParseEAN13 takes n EAN13 code and parses it
func ParseEAN13(code string) (EAN13, error) {
	ean := EAN13{Prefix: gs1.PrefixUnknown}

	if len(code) != 13 {
		return ean, fmt.Errorf("EAN13: code must have 13 digits")
	}

	if _, err := strconv.Atoi(code); err != nil {
		return ean, fmt.Errorf("EAN13: code must only contain digits")
	}

	ean.Data = code[:12]
	ean.Parity = int(code[12] - '0')
	parityExpected := gs1.CalcParity(ean.Data)

	if ean.Parity != parityExpected {
		return ean, fmt.Errorf("EAN13: parity isn't valid. Got %d, should %d. Is there a typo in the code?", ean.Parity, parityExpected)
	}

	if strings.HasPrefix(ean.Data, "0000000") {
		ean.Prefix = gs1.PrefixReserved
		return ean, fmt.Errorf("EAN13: code has reserved prefix, which is not valid for current usage")
	}

	threeDigit, err := strconv.Atoi(ean.Data[:3])
	if err != nil {
		panic(err)
	}
	ean.Prefix = gs1.CountryCode(threeDigit)
	ean.Data = ean.Data[3:]

	if ean.Prefix == gs1.PrefixReserved || ean.Prefix == gs1.PrefixUSReserved {
		return ean, fmt.Errorf("EAN13: code has reserved prefix, which is not valid for current usage")
	}

	return ean, nil
}
