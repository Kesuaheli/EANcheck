package gs1

// Prefix represents the country prefix of a number code
type Prefix string

// Country code
const (
	// Used to issue Restricted Circulation Numbers within a company
	PrefixCompany Prefix = "Within company"
	// Used to issue Restricted Circulation Numbers within a geographic region
	PrefixRegion Prefix = "Within geografic region"
	// Reserved by GS1 Global Office for future use
	PrefixReserved Prefix = "Reserved"
	// Reserved by GS1 US for future use
	PrefixUSReserved Prefix = "US, Reserved"
	// Used when an parsing error occurs
	PrefixUnknown Prefix = "Unknown"

	//
	PrefixISSN Prefix = "Serial publications (ISSN)"
	//
	PrefixISBN Prefix = "Bookland (ISBN)"
	//
	PrefixRefundReceipts Prefix = "Refund receipts"
	//
	PrefixCouponIdentification Prefix = "Coupon identification"

	PrefixUSCanada            Prefix = "US, Canada"
	PrefixUS                  Prefix = "US"
	PrefixFrance              Prefix = "France"
	PrefixBulgaria            Prefix = "Bulgaria"
	PrefixSlovenija           Prefix = "Slovenija"
	PrefixCroatia             Prefix = "Croatia"
	PrefixBHI                 Prefix = "Bosnia-Herzegovina"
	PrefixMontenegro          Prefix = "Montenegro"
	PrefixGermany             Prefix = "Germany"
	PrefixJapan               Prefix = "Japan"
	PrefixRussia              Prefix = "Russia"
	PrefixKyrgyzstan          Prefix = "Kyrgyzstan"
	PrefixChineseTaipei       Prefix = "Chinese Taipei"
	PrefixEstonia             Prefix = "Estonia"
	PrefixLatvia              Prefix = "Latvia"
	PrefixAzerbaijan          Prefix = "Azerbaijan"
	PrefixLithuania           Prefix = "Lithuania"
	PrefixUzbekistan          Prefix = "Uzbekistan"
	PrefixSriLanka            Prefix = "Sri Lanka"
	PrefixPhilippines         Prefix = "Philippines"
	PrefixBelarus             Prefix = "Belarus"
	PrefixUkraine             Prefix = "Ukraine"
	PrefixTurkmenistan        Prefix = "Turkmenistan"
	PrefixMoldova             Prefix = "Moldova"
	PrefixArmenia             Prefix = "Armenia"
	PrefixGeorgia             Prefix = "Georgia"
	PrefixKazakstan           Prefix = "Kazakstan"
	PrefixTajikistan          Prefix = "Tajikistan"
	PrefixChinaHongKong       Prefix = "Hong Kong, China"
	PrefixUK                  Prefix = "UK"
	PrefixGreece              Prefix = "Association Greece"
	PrefixLebanon             Prefix = "Lebanon"
	PrefixCyprus              Prefix = "Cyprus"
	PrefixAlbania             Prefix = "Albania"
	PrefixMacedonia           Prefix = "Macedonia"
	PrefixMalta               Prefix = "Malta"
	PrefixIreland             Prefix = "Ireland"
	PrefixBelgiumLuxembourg   Prefix = "Belgium & Luxembourg"
	PrefixPortugal            Prefix = "Portugal"
	PrefixIceland             Prefix = "Iceland"
	PrefixDenmark             Prefix = "Denmark"
	PrefixPoland              Prefix = "Poland"
	PrefixRomania             Prefix = "Romania"
	PrefixHungary             Prefix = "Hungary"
	PrefixSouthAfrica         Prefix = "South Africa"
	PrefixGhana               Prefix = "Ghana"
	PrefixSenegal             Prefix = "Senegal"
	PrefixOman                Prefix = "Oman"
	PrefixBahrain             Prefix = "Bahrain"
	PrefixMauritius           Prefix = "Mauritius"
	PrefixMorocco             Prefix = "Morocco"
	PrefixAlgeria             Prefix = "Algeria"
	PrefixNigeria             Prefix = "Nigeria"
	PrefixKenya               Prefix = "Kenya"
	PrefixCameroon            Prefix = "Cameroon"
	PrefixCotedIvoire         Prefix = "Côte d'Ivoire"
	PrefixTunisia             Prefix = "Tunisia"
	PrefixTanzania            Prefix = "Tanzania"
	PrefixSyria               Prefix = "Syria"
	PrefixEgypt               Prefix = "Egypt"
	PrefixLibya               Prefix = "Libya"
	PrefixJordan              Prefix = "Jordan"
	PrefixIran                Prefix = "Iran"
	PrefixKuwait              Prefix = "Kuwait"
	PrefixSaudiArabia         Prefix = "Saudi Arabia"
	PrefixEmirates            Prefix = "Emirates"
	PrefixQatar               Prefix = "Qatar"
	PrefixNamibia             Prefix = "Namibia"
	PrefixFinland             Prefix = "Finland"
	PrefixChina               Prefix = "China"
	PrefixNorway              Prefix = "Norway"
	PrefixIsrael              Prefix = "Israel"
	PrefixSweden              Prefix = "Sweden"
	PrefixGuatemala           Prefix = "Guatemala"
	PrefixElSalvador          Prefix = "El Salvador"
	PrefixHonduras            Prefix = "Honduras"
	PrefixNicaragua           Prefix = "Nicaragua"
	PrefixCostaRica           Prefix = "Costa Rica"
	PrefixPanama              Prefix = "Panama"
	PrefixRepublicaDominicana Prefix = "Republica Dominicana"
	PrefixMexico              Prefix = "Mexico"
	PrefixCanada              Prefix = "Canada"
	PrefixVenezuela           Prefix = "Venezuela"
	PrefixSwitzerland         Prefix = "Switzerland"
	PrefixColombia            Prefix = "Colombia"
	PrefixUruguay             Prefix = "Uruguay"
	PrefixPeru                Prefix = "Peru"
	PrefixBolivia             Prefix = "Bolivia"
	PrefixArgentina           Prefix = "Argentina"
	PrefixChile               Prefix = "Chile"
	PrefixParaguay            Prefix = "Paraguay"
	PrefixEcuador             Prefix = "Ecuador"
	PrefixBrasil              Prefix = "Brasil"
	PrefixItaly               Prefix = "Italy"
	PrefixSpain               Prefix = "Spain"
	PrefixCuba                Prefix = "Cuba"
	PrefixSlovakia            Prefix = "Slovakia"
	PrefixCzech               Prefix = "Czech"
	PrefixSerbia              Prefix = "Serbia"
	PrefixMongolia            Prefix = "Mongolia"
	PrefixNorthKorea          Prefix = "North Korea"
	PrefixTurkiye             Prefix = "Türkiye"
	PrefixNetherlands         Prefix = "Netherlands"
	PrefixSouthKorea          Prefix = "South Korea"
	PrefixMyanmar             Prefix = "Myanmar"
	PrefixCambodia            Prefix = "Cambodia"
	PrefixThailand            Prefix = "Thailand"
	PrefixSingapore           Prefix = "Singapore"
	PrefixIndia               Prefix = "India"
	PrefixVietnam             Prefix = "Vietnam"
	PrefixPakistan            Prefix = "Pakistan"
	PrefixIndonesia           Prefix = "Indonesia"
	PrefixAustria             Prefix = "Austria"
	PrefixAustralia           Prefix = "Australia"
	PrefixNewZealand          Prefix = "New Zealand"
	PrefixMalaysia            Prefix = "Malaysia"
	PrefixChinaMacao          Prefix = "Macao, China"
)

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
