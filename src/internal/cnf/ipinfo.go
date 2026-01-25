package cnf

import "net"

// DBTypeIPInfo identifies IPInfo as the database provider.
const DBTypeIPInfo DBType = 1

// IPInfo schema: https://github.com/ipinfo/sample-database/

// IPInfoCountry represents country-level GeoIP data from IPInfo.
type IPInfoCountry struct {
	StartIP       net.IP `maxminddb:"start_ip"`
	EndIP         net.IP `maxminddb:"end_ip"`
	Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`
}

// IPInfoASN represents ASN data from IPInfo.
type IPInfoASN struct {
	StartIP net.IP `maxminddb:"start_ip"`
	EndIP   net.IP `maxminddb:"end_ip"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
}

// IPInfoASNExtended represents extended ASN data from IPInfo.
type IPInfoASNExtended struct {
	StartIP net.IP `maxminddb:"start_ip"`
	EndIP   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
	Type    string `maxminddb:"type"`
	Country string `maxminddb:"country"`
}

// IPInfoPrivacy represents privacy detection data from IPInfo.
type IPInfoPrivacy struct {
	StartIP net.IP `maxminddb:"start_ip"`
	EndIP   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	Hosting bool   `maxminddb:"hosting"`
	Proxy   bool   `maxminddb:"proxy"`
	Tor     bool   `maxminddb:"tor"`
	VPN     bool   `maxminddb:"vpn"`
	Relay   bool   `maxminddb:"relay"`
	Service string `maxminddb:"service"`
}

// IPInfoCity represents city-level GeoIP data from IPInfo.
type IPInfoCity struct {
	StartIP    net.IP  `maxminddb:"start_ip"`
	EndIP      net.IP  `maxminddb:"end_ip"`
	JoinKey    net.IP  `maxminddb:"join_key"`
	City       string  `maxminddb:"city"`
	Region     string  `maxminddb:"region"`
	Country    string  `maxminddb:"country"`
	Latitude   float32 `maxminddb:"latitude"`
	Longitude  float32 `maxminddb:"longitude"`
	PostalCode string  `maxminddb:"postal_code"`
	Timezone   string  `maxminddb:"timezone"`
}

// TODO: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Company
// TODO: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Mobile%20Carrier
// TODO: https://github.com/ipinfo/sample-database/tree/main/IP%20Geolocation%20Extended
// TODO: https://github.com/ipinfo/sample-database/tree/main/Privacy%20Detection%20Extended
// TODO: https://github.com/ipinfo/sample-database/tree/main/Abuse%20Contact
// TODO: https://github.com/ipinfo/sample-database/tree/main/Hosted%20Domains
