package lookup

import (
	"net"

	"github.com/tracyhatemice/gogeoip-lookup/src/internal/cnf"
)

// IPInfoCountry looks up country information for an IP using IPInfo database.
func IPInfoCountry(ip net.IP) (any, error) {
	return lookupGeneric[cnf.IPInfoCountry](ip, cnf.DBCountry)
}

// IPInfoCity looks up city information for an IP using IPInfo database.
func IPInfoCity(ip net.IP) (any, error) {
	return lookupGeneric[cnf.IPInfoCity](ip, cnf.DBCity)
}

// IPInfoASN looks up ASN information for an IP using IPInfo database.
func IPInfoASN(ip net.IP) (any, error) {
	return lookupGeneric[cnf.IPInfoASN](ip, cnf.DBASN)
}

// IPInfoPrivacy looks up privacy detection information using IPInfo database.
func IPInfoPrivacy(ip net.IP) (any, error) {
	return lookupGeneric[cnf.IPInfoPrivacy](ip, cnf.DBPrivacy)
}
