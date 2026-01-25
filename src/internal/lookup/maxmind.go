package lookup

import (
	"net"

	"github.com/tracyhatemice/gogeoip-lookup/src/internal/cnf"
)

// MaxMindCountry looks up country information for an IP using MaxMind database.
func MaxMindCountry(ip net.IP) (any, error) {
	return lookupGeneric[cnf.MaxMindCountry](ip, cnf.DBCountry)
}

// MaxMindCity looks up city information for an IP using MaxMind database.
func MaxMindCity(ip net.IP) (any, error) {
	return lookupGeneric[cnf.MaxMindCity](ip, cnf.DBCity)
}

// MaxMindASN looks up ASN information for an IP using MaxMind database.
func MaxMindASN(ip net.IP) (any, error) {
	return lookupGeneric[cnf.MaxMindASN](ip, cnf.DBASN)
}

// MaxMindPrivacy looks up privacy/anonymous IP information using MaxMind database.
func MaxMindPrivacy(ip net.IP) (any, error) {
	return lookupGeneric[cnf.MaxMindPrivacy](ip, cnf.DBPrivacy)
}
