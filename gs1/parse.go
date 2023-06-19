package gs1

import (
	"strconv"
	"strings"
)

// ParseEAN13 takes n EAN13 code and parses it
func ParseEAN13(code string) (prefix Prefix, rest string) {
	if strings.HasPrefix(code, "0000000") {
		return PrefixReserved, ""
	}

	threeDigit, err := strconv.Atoi(code[:3])
	if err != nil {
		panic(err)
	}

	rest = code[3:]
	prefix = CountryCode(threeDigit)
	if prefix == PrefixReserved || prefix == PrefixUSReserved {
		rest = ""
	}
	return prefix, rest
}

// CountryCode takes a three digit integer and converts it to
func CountryCode(code int) Prefix {
	// test if code is in range of a and b (inclusive)
	r := func(a, b int) bool {
		return code >= a && code <= b
	}

	switch {
	case r(0, 19):
		return PrefixUSCanada
	case r(20, 29):
		return PrefixRegion
	case r(30, 39):
		return PrefixUS
	case r(40, 49):
		return PrefixCompany
	case r(50, 59):
		return PrefixUSReserved
	case r(60, 99):
		return PrefixUSCanada
	case r(100, 139):
		return PrefixUS
	case r(200, 299):
		return PrefixRegion

	case r(400, 440):
		return PrefixGermany

		//TODO: other ranges
	}
	return PrefixReserved
}
