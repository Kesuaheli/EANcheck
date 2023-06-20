package gs1

// Prefix represents the country prefix of a number code
type Prefix string

// EAN13 represents an EAN13 code
type EAN13 struct {
	Prefix Prefix
	Data   string
	Parity int
}

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
