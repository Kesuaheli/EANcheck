package gs1

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseEAN13 takes n EAN13 code and parses it
func ParseEAN13(code string) (EAN13, error) {
	ean := EAN13{Prefix: PrefixUnknown}

	if len(code) != 13 {
		return ean, fmt.Errorf("EAN13: code must have 13 digits")
	}

	if _, err := strconv.Atoi(code); err != nil {
		return ean, fmt.Errorf("EAN13: code must only contain digits")
	}

	ean.Data = code[:12]
	ean.Parity = int(code[12] - '0')
	parityExpected := CalcParity(ean.Data)

	if ean.Parity != parityExpected {
		return ean, fmt.Errorf("EAN13: parity isn't valid. Got %d, should %d. Is there a typo in the code?", ean.Parity, parityExpected)
	}

	if strings.HasPrefix(ean.Data, "0000000") {
		ean.Prefix = PrefixReserved
		return ean, fmt.Errorf("EAN13: code has reserved prefix, which is not valid for current usage")
	}

	threeDigit, err := strconv.Atoi(ean.Data[:3])
	if err != nil {
		panic(err)
	}
	ean.Prefix = CountryCode(threeDigit)
	ean.Data = ean.Data[3:]

	if ean.Prefix == PrefixReserved || ean.Prefix == PrefixUSReserved {
		return ean, fmt.Errorf("EAN13: code has reserved prefix, which is not valid for current usage")
	}

	return ean, nil
}

// CountryCode takes a three digit integer and returns it Prefix.
// Source: https://www.gs1.org/standards/id-keys/company-prefix
func CountryCode(code int) Prefix {
	// s tests if code matches a single integer
	var s func(int) bool = func(a int) bool {
		return code == a
	}
	// r tests if code is in range of a and b (inclusive)
	var r func(int, int) bool = func(a, b int) bool {
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
	case r(300, 379):
		return PrefixFrance
	case s(380):
		return PrefixBulgaria
	case s(383):
		return PrefixSlovenija
	case s(385):
		return PrefixCroatia
	case s(387):
		return PrefixBHI
	case s(389):
		return PrefixMontenegro
	case r(400, 440):
		return PrefixGermany
	case r(450, 459):
		return PrefixJapan
	case r(460, 469):
		return PrefixRussia
	case s(470):
		return PrefixKyrgyzstan
	case s(471):
		return PrefixChineseTaipei
	case s(474):
		return PrefixEstonia
	case s(475):
		return PrefixLatvia
	case s(476):
		return PrefixAzerbaijan
	case s(477):
		return PrefixLithuania
	case s(478):
		return PrefixUzbekistan
	case s(479):
		return PrefixSriLanka
	case s(480):
		return PrefixPhilippines
	case s(481):
		return PrefixBelarus
	case s(482):
		return PrefixUkraine
	case s(483):
		return PrefixTurkmenistan
	case s(484):
		return PrefixMoldova
	case s(485):
		return PrefixArmenia
	case s(486):
		return PrefixGeorgia
	case s(487):
		return PrefixKazakstan
	case s(488):
		return PrefixTajikistan
	case s(489):
		return PrefixChinaHongKong
	case r(490, 499):
		return PrefixJapan
	case r(500, 509):
		return PrefixUK
	case r(520, 521):
		return PrefixGreece
	case s(528):
		return PrefixLebanon
	case s(529):
		return PrefixCyprus
	case s(530):
		return PrefixAlbania
	case s(531):
		return PrefixMacedonia
	case s(535):
		return PrefixMalta
	case s(539):
		return PrefixIreland
	case r(540, 549):
		return PrefixBelgiumLuxembourg
	case s(560):
		return PrefixPortugal
	case s(569):
		return PrefixIceland
	case r(570, 579):
		return PrefixDenmark
	case s(590):
		return PrefixPoland
	case s(594):
		return PrefixRomania
	case s(599):
		return PrefixHungary
	case r(600, 601):
		return PrefixSouthAfrica
	case s(603):
		return PrefixGhana
	case s(604):
		return PrefixSenegal
	case r(605, 606):
		return PrefixReserved
	case s(607):
		return PrefixOman
	case s(608):
		return PrefixBahrain
	case s(609):
		return PrefixMauritius
	case s(610):
		return PrefixReserved
	case s(611):
		return PrefixMorocco
	case s(613):
		return PrefixAlgeria
	case s(614):
		return PrefixReserved
	case s(615):
		return PrefixNigeria
	case s(616):
		return PrefixKenya
	case s(617):
		return PrefixCameroon
	case s(618):
		return PrefixCotedIvoire
	case s(619):
		return PrefixTunisia
	case s(620):
		return PrefixTanzania
	case s(621):
		return PrefixSyria
	case s(622):
		return PrefixEgypt
	case s(623):
		return PrefixReserved
	case s(624):
		return PrefixLibya
	case s(625):
		return PrefixJordan
	case s(626):
		return PrefixIran
	case s(627):
		return PrefixKuwait
	case s(628):
		return PrefixSaudiArabia
	case s(629):
		return PrefixEmirates
	case s(630):
		return PrefixQatar
	case s(631):
		return PrefixNamibia
	case r(640, 649):
		return PrefixFinland
	case r(690, 699):
		return PrefixChina
	case r(700, 709):
		return PrefixNorway
	case s(729):
		return PrefixIsrael
	case r(730, 739):
		return PrefixSweden
	case s(740):
		return PrefixGuatemala
	case s(741):
		return PrefixElSalvador
	case s(742):
		return PrefixHonduras
	case s(743):
		return PrefixNicaragua
	case s(744):
		return PrefixCostaRica
	case s(745):
		return PrefixPanama
	case s(746):
		return PrefixRepublicaDominicana
	case s(750):
		return PrefixMexico
	case r(754, 755):
		return PrefixCanada
	case s(758):
		return PrefixReserved
	case s(759):
		return PrefixVenezuela
	case r(760, 769):
		return PrefixSwitzerland
	case r(770, 771):
		return PrefixColombia
	case s(773):
		return PrefixUruguay
	case s(775):
		return PrefixPeru
	case s(777):
		return PrefixBolivia
	case r(778, 779):
		return PrefixArgentina
	case s(780):
		return PrefixChile
	case s(784):
		return PrefixParaguay
	case s(786):
		return PrefixEcuador
	case r(789, 790):
		return PrefixBrasil
	case r(800, 839):
		return PrefixItaly
	case r(840, 849):
		return PrefixSpain
	case s(850):
		return PrefixCuba
	case s(858):
		return PrefixSlovakia
	case s(859):
		return PrefixCzech
	case s(860):
		return PrefixSerbia
	case s(865):
		return PrefixMongolia
	case s(867):
		return PrefixNorthKorea
	case r(868, 869):
		return PrefixTurkiye
	case r(870, 879):
		return PrefixNetherlands
	case s(880):
		return PrefixSouthKorea
	case s(883):
		return PrefixMyanmar
	case s(884):
		return PrefixCambodia
	case s(885):
		return PrefixThailand
	case s(888):
		return PrefixSingapore
	case s(890):
		return PrefixIndia
	case s(893):
		return PrefixVietnam
	case s(894):
		return PrefixReserved
	case s(896):
		return PrefixPakistan
	case s(899):
		return PrefixIndonesia
	case r(900, 919):
		return PrefixAustria
	case r(930, 939):
		return PrefixAustralia
	case r(940, 949):
		return PrefixNewZealand
	case r(950, 952):
		return PrefixReserved
	case s(955):
		return PrefixMalaysia
	case s(958):
		return PrefixChinaMacao
	case r(960, 969):
		return PrefixReserved
	case s(977):
		return PrefixISSN
	case r(978, 979):
		return PrefixISBN
	case s(980):
		return PrefixRefundReceipts
	case r(981, 983):
		return PrefixCouponIdentification
	case r(990, 999):
		return PrefixCouponIdentification
	}

	// https://www.gs1.org/standards/id-keys/company-prefix => see Note 1
	// "Prefixes not explicitly listed above are reserved by GS1 Global Office for future use"
	return PrefixReserved
}

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
