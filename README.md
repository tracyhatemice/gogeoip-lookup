[![Lint](https://github.com/tracyhatemice/gogeoip-lookup/actions/workflows/lint.yml/badge.svg?branch=latest)](https://github.com/tracyhatemice/gogeoip-lookup/actions/workflows/lint.yml)

# GeoIP Lookup Service

Go-based microservice to perform IP lookups in local GeoIP databases.

It currently only supports databases in **MMDB format**.

If you want to use their extended databases, you might encounter problems. You are welcome to help integrating them correctly.

Feel free to [open a ticket](https://github.com/tracyhatemice/gogeoip-lookup/issues/new) if you encounter any issues.

----

## GeoIP Provider Support

* **IPInfo**: [Information](https://ipinfo.io/products/free-ip-database), [CC4 License](https://creativecommons.org/licenses/by-sa/4.0/) (*allows for commercial usage - you need to add an attribution*)

    **Attribution**: `<p>IP address data powered by <a href="https://ipinfo.io">IPinfo</a></p>`

* **MaxMind**: [Information](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data), [EULA](https://www.maxmind.com/en/geolite2/eula) (*allows for limited commercial usage - you need to add an attribution*)

    **Attribution**: `This product includes GeoLite2 data created by MaxMind, available from <a href="https://www.maxmind.com">https://www.maxmind.com</a>.`

These two providers were tested.

----

## Integration

* [HAProxy Community using Lua](https://github.com/O-X-L/haproxy-geoip)

   NOTE: HAProxy provides enterprise-grade licensing that has this functionality built-in.

Make sure to read the GeoIP-DB License before integrating it with any service!

----

## Usage

The binary starts an HTTP server with configurable timeouts.

### Command Line Options

```bash
./geoip-lookup [options]

Options:
  -l        Listen address (default: 127.0.0.1)
  -p        Listen port (default: 10000)
  -t        Database type: ipinfo or maxmind (default: ipinfo)
  -country  Path to country database (default: /etc/geoip/country.mmdb)
  -city     Path to city database (default: /etc/geoip/city.mmdb)
  -asn      Path to ASN database (default: /etc/geoip/asn.mmdb)
  -privacy  Path to privacy database (default: /etc/geoip/privacy.mmdb)
  -plain    Return plain text instead of JSON (default: false)
```

### API Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /lookup/{type}` | Perform GeoIP lookup |
| `GET /health` | Health check endpoint |

**Lookup types:** `country`, `city`, `asn`, `privacy`, `country_asn` (IPInfo only)

**Query parameters:**
- `ip` - IP address to lookup (optional, defaults to client IP)
- `filter` - Dot-separated path to extract specific fields

### Examples

```bash
# Start the service
./geoip-lookup -l 127.0.0.1 -p 10069 -t ipinfo \
  -country /etc/geoip/country.mmdb \
  -asn /etc/geoip/asn.mmdb \
  -city /etc/geoip/city.mmdb

# IPInfo examples
curl "http://127.0.0.1:10069/lookup/country?ip=1.1.1.1"
> {"Country":"US","CountryName":"United States","Continent":"NA","ContinentName":"North America",...}

curl "http://127.0.0.1:10069/lookup/country?ip=1.1.1.1&filter=Country"
> "US"

curl "http://127.0.0.1:10069/lookup/asn?ip=1.1.1.1"
> {"ASN":"AS13335","Name":"Cloudflare, Inc.","Domain":"cloudflare.com"}

# MaxMind examples
./geoip-lookup -t maxmind ...

curl "http://127.0.0.1:10069/lookup/asn?ip=1.1.1.1"
> {"ASN":"13335","Name":"CLOUDFLARENET"}

curl "http://127.0.0.1:10069/lookup/country?ip=8.8.8.8"
> {"Country":{"Code":"US","ID":6252001,"EuropeanUnion":false},"Continent":{"Code":"NA",...},...}

# Filter nested fields
curl "http://127.0.0.1:10069/lookup/country?ip=8.8.8.8&filter=Country.Code"
> "US"

curl "http://127.0.0.1:10069/lookup/city?ip=8.8.8.8&filter=Location"
> {"AccuracyRadius":1000,"Latitude":37.751,"Longitude":-97.822,"Timezone":"America/Chicago"}

# Health check
curl "http://127.0.0.1:10069/health"
> {"status":"ok"}

# Plain text output
./geoip-lookup -plain ...
curl "http://127.0.0.1:10069/lookup/country?ip=1.1.1.1&filter=CountryName"
> United States

# Listen on all interfaces
./geoip-lookup -l 0.0.0.0 -p 10069 ...
```

----

## Testing

Basic integration tests are done by using the test-script:

```bash
bash scripts/test.sh
```

Feel free to contribute more test-cases if you found some edge-case issue(s).

----

## Service

Example systemd service:

```text
[Unit]
Description=GeoIP Lookup Service
Documentation=https://github.com/tracyhatemice/gogeoip-lookup

[Service]
Type=simple
ExecStart=/usr/bin/geoip-lookup -l 127.0.0.1 -p 10069 -t ipinfo -country /etc/geoip/country.mmdb -asn /etc/geoip/asn.mmdb -city /etc/geoip/city.mmdb

# service-user only needs read-access to databases
User=geoip
Group=geoip
Restart=on-failure
RestartSec=5s

StandardOutput=journal
StandardError=journal
SyslogIdentifier=geoip-lookup

[Install]
WantedBy=multi-user.target
```

----
## Thanks

O-X-L/geoip-lookup-service
