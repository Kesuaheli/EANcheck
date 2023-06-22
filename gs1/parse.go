package gs1

// CalcParity takes a GS1 code and calculates an returns its parity number.
// This funtion only takes the actual data digits an not the last digit.
//
// Source: https://www.gs1.org/services/how-calculate-check-digit-manually
func CalcParity(code string) (parity int) {
	codeB := []byte(code)

	//reverse slice
	for i, j := 0, len(codeB)-1; i < j; i, j = i+1, j-1 {
		codeB[i], codeB[j] = codeB[j], codeB[i]
	}

	// weighted checksum
	for i, d := range codeB {
		weight := (i+1)%2*2 + 1
		parity += int(d-'0') * weight
	}

	// checksum to parity
	parity = (10 - parity%10) % 10

	return parity
}
