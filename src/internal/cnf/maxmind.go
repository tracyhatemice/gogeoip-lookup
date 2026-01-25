package cnf

// DBTypeMaxMind identifies MaxMind as the database provider.
const DBTypeMaxMind DBType = 2

// MaxMind schema: https://github.com/maxmind/MaxMind-DB/tree/main/source-data

// MaxMindCountryInfo contains nested types for country data.
type MaxMindCountryInfo struct {
	Code          string            `maxminddb:"iso_code"`
	ID            uint              `maxminddb:"geoname_id"`
	Names         map[string]string `maxminddb:"names"`
	EuropeanUnion bool              `maxminddb:"is_in_european_union"`
}

// MaxMindContinentInfo contains continent data.
type MaxMindContinentInfo struct {
	Code  string            `maxminddb:"code"`
	ID    uint              `maxminddb:"geoname_id"`
	Names map[string]string `maxminddb:"names"`
}

// MaxMindCountry represents country-level GeoIP data from MaxMind.
type MaxMindCountry struct {
	Country           MaxMindCountryInfo   `maxminddb:"country"`
	RegisteredCountry MaxMindCountryInfo   `maxminddb:"registered_country"`
	Continent         MaxMindContinentInfo `maxminddb:"continent"`
}

// MaxMindASN represents ASN data from MaxMind.
type MaxMindASN struct {
	ASN  uint   `maxminddb:"autonomous_system_number"`
	Name string `maxminddb:"autonomous_system_organization"`
}

// MaxMindCityInfo contains city-specific data.
type MaxMindCityInfo struct {
	ID    uint              `maxminddb:"geoname_id"`
	Names map[string]string `maxminddb:"names"`
}

// MaxMindLocationInfo contains geographic location data.
type MaxMindLocationInfo struct {
	AccuracyRadius uint    `maxminddb:"accuracy_radius"`
	Latitude       float32 `maxminddb:"latitude"`
	Longitude      float32 `maxminddb:"longitude"`
	Timezone       string  `maxminddb:"time_zone"`
}

// MaxMindPostalInfo contains postal code data.
type MaxMindPostalInfo struct {
	Code string `maxminddb:"code"`
}

// MaxMindTraitsInfo contains IP traits data.
type MaxMindTraitsInfo struct {
	IsAnycast        bool `maxminddb:"is_anycast"`
	IsAnonymousProxy bool `maxminddb:"is_anonymous_proxy"`
}

// MaxMindCity represents city-level GeoIP data from MaxMind.
type MaxMindCity struct {
	City              MaxMindCityInfo      `maxminddb:"city"`
	Country           MaxMindCountryInfo   `maxminddb:"country"`
	RegisteredCountry MaxMindCountryInfo   `maxminddb:"registered_country"`
	Continent         MaxMindContinentInfo `maxminddb:"continent"`
	Location          MaxMindLocationInfo  `maxminddb:"location"`
	Postal            MaxMindPostalInfo    `maxminddb:"postal"`
	Traits            MaxMindTraitsInfo    `maxminddb:"traits"`
}

// MaxMindPrivacy represents privacy/anonymous IP data from MaxMind.
type MaxMindPrivacy struct {
	Any          bool `maxminddb:"is_anonymous"`
	VPN          bool `maxminddb:"is_anonymous_vpn"`
	Tor          bool `maxminddb:"is_tor_exit_node"`
	Hosting      bool `maxminddb:"is_hosting_provider"`
	PublicProxy  bool `maxminddb:"is_public_proxy"`
	PrivateProxy bool `maxminddb:"is_residential_proxy"`
}
