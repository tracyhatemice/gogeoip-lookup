package lookup

import (
	"errors"
	"net"
	"net/netip"

	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/tracyhatemice/gogeoip-lookup/src/internal/cnf"
)

// LookupFunc defines the signature for GeoIP lookup functions.
type LookupFunc func(ip net.IP) (any, error)

// FuncMapping maps database types to their lookup function implementations.
var FuncMapping = map[cnf.DBType]map[string]LookupFunc{
	cnf.DBTypeIPInfo: {
		"country": IPInfoCountry,
		"city":    IPInfoCity,
		"asn":     IPInfoASN,
		"privacy": IPInfoPrivacy,
	},
	cnf.DBTypeMaxMind: {
		"country": MaxMindCountry,
		"city":    MaxMindCity,
		"asn":     MaxMindASN,
		"privacy": MaxMindPrivacy,
	},
}

// Funcs returns the lookup functions for the current database type.
func Funcs() map[string]LookupFunc {
	return FuncMapping[cnf.CurrentDBType]
}

// toNetipAddr converts a net.IP to netip.Addr for the v2 API.
func toNetipAddr(ip net.IP) (netip.Addr, error) {
	addr, ok := netip.AddrFromSlice(ip)
	if !ok {
		return netip.Addr{}, errors.New("invalid IP address")
	}
	return addr, nil
}

func lookupGeneric[T any](ip net.IP, dbFile string) (*T, error) {
	db, err := maxminddb.Open(dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	addr, err := toNetipAddr(ip)
	if err != nil {
		return nil, err
	}

	var result T
	err = db.Lookup(addr).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
