package cnf

// DBType represents the type of GeoIP database being used.
type DBType uint8

// Configuration variables for the GeoIP lookup service.
var (
	CurrentDBType = DBTypeIPInfo
	DBCountry     = "/etc/geoip/country.mmdb"
	DBCity        = "/etc/geoip/city.mmdb"
	DBASN         = "/etc/geoip/asn.mmdb"
	DBPrivacy     = "/etc/geoip/privacy.mmdb" // vpns, tor, proxies, hosting
	ReturnPlain   = false
)
