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

## Quick Start with Docker

### Using Pre-built Image

```bash
docker pull ghcr.io/tracyhatemice/gogeoip-lookup:latest
```

### Using Docker Compose (MaxMind)

The `contrib/docker-maxmind/` directory contains a complete setup with automatic database updates.

1. Create a `geoip.env` file with your MaxMind credentials:
   ```bash
   GEOIPUPDATE_ACCOUNT_ID=your_account_id
   GEOIPUPDATE_LICENSE_KEY=your_license_key
   ```

2. Start the services:
   ```bash
   cd contrib/docker-maxmind
   docker compose up -d
   ```

This will:
- Start the GeoIP lookup service on port `80`
- Automatically download and update MaxMind GeoLite2 databases every 72 hours

### Container Configuration

| Option | Description | Default |
|--------|-------------|---------|
| `-l` | Listen address (dual-stack) | `::` |
| `-p` | Listen port | `80` |
| `-t` | Database type (`ipinfo` or `maxmind`) | `ipinfo` |
| `-country` | Path to country database | `/etc/geoip/country.mmdb` |
| `-city` | Path to city database | `/etc/geoip/city.mmdb` |
| `-asn` | Path to ASN database | `/etc/geoip/asn.mmdb` |
| `-privacy` | Path to privacy database | `/etc/geoip/privacy.mmdb` |
| `-plain` | Return plain text instead of JSON | `false` |

----

## Integration

* [HAProxy Community using Lua](https://github.com/O-X-L/haproxy-geoip)

   NOTE: HAProxy provides enterprise-grade licensing that has this functionality built-in.

Make sure to read the GeoIP-DB License before integrating it with any service!

----

## API Reference

### Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /lookup/{type}` | Perform GeoIP lookup |
| `GET /health` | Health check endpoint |

**Lookup types:** `country`, `city`, `asn`, `privacy`

**Query parameters:**
- `ip` - IP address to lookup (optional, defaults to client IP)
- `filter` - Dot-separated path to extract specific fields

### Response Fields & Filters

#### IPInfo Database

| Lookup | Field | Type | Filter Example |
|--------|-------|------|----------------|
| **country** | `Country` | string | `filter=Country` |
| | `CountryName` | string | `filter=CountryName` |
| | `Continent` | string | `filter=Continent` |
| | `ContinentName` | string | `filter=ContinentName` |
| **city** | `City` | string | `filter=City` |
| | `Region` | string | `filter=Region` |
| | `Country` | string | `filter=Country` |
| | `Latitude` | float | `filter=Latitude` |
| | `Longitude` | float | `filter=Longitude` |
| | `PostalCode` | string | `filter=PostalCode` |
| | `Timezone` | string | `filter=Timezone` |
| **asn** | `ASN` | string | `filter=ASN` |
| | `Name` | string | `filter=Name` |
| | `Domain` | string | `filter=Domain` |
| **privacy** | `Hosting` | bool | `filter=Hosting` |
| | `Proxy` | bool | `filter=Proxy` |
| | `Tor` | bool | `filter=Tor` |
| | `VPN` | bool | `filter=VPN` |
| | `Relay` | bool | `filter=Relay` |
| | `Service` | string | `filter=Service` |

#### MaxMind Database

| Lookup | Field | Type | Filter Example |
|--------|-------|------|----------------|
| **country** | `Country.Code` | string | `filter=Country.Code` |
| | `Country.ID` | uint | `filter=Country.ID` |
| | `Country.EuropeanUnion` | bool | `filter=Country.EuropeanUnion` |
| | `RegisteredCountry.Code` | string | `filter=RegisteredCountry.Code` |
| | `RegisteredCountry.ID` | uint | `filter=RegisteredCountry.ID` |
| | `Continent.Code` | string | `filter=Continent.Code` |
| | `Continent.ID` | uint | `filter=Continent.ID` |
| | `Continent.Names` | map | `filter=Continent.Names` |
| **city** | `City.Code` | string | `filter=City.Code` |
| | `City.ID` | uint | `filter=City.ID` |
| | `Country.Code` | string | `filter=Country.Code` |
| | `Location.Latitude` | float | `filter=Location.Latitude` |
| | `Location.Longitude` | float | `filter=Location.Longitude` |
| | `Location.Timezone` | string | `filter=Location.Timezone` |
| | `Location.AccuracyRadius` | uint | `filter=Location.AccuracyRadius` |
| | `Postal.Code` | string | `filter=Postal.Code` |
| | `Traits.IsAnycast` | bool | `filter=Traits.IsAnycast` |
| | `Traits.IsAnonymousProxy` | bool | `filter=Traits.IsAnonymousProxy` |
| **asn** | `ASN` | string | `filter=ASN` |
| | `Name` | string | `filter=Name` |
| **privacy** | `Any` | bool | `filter=Any` |
| | `VPN` | bool | `filter=VPN` |
| | `Tor` | bool | `filter=Tor` |
| | `Hosting` | bool | `filter=Hosting` |
| | `PublicProxy` | bool | `filter=PublicProxy` |
| | `PrivateProxy` | bool | `filter=PrivateProxy` |

### Examples

```bash
# IPInfo examples
curl "http://localhost:10069/lookup/country?ip=1.1.1.1"
> {"Country":"US","CountryName":"United States","Continent":"NA","ContinentName":"North America",...}

curl "http://localhost:10069/lookup/country?ip=1.1.1.1&filter=Country"
> "US"

curl "http://localhost:10069/lookup/asn?ip=1.1.1.1"
> {"ASN":"AS13335","Name":"Cloudflare, Inc.","Domain":"cloudflare.com"}

# MaxMind examples
curl "http://localhost:10069/lookup/asn?ip=1.1.1.1"
> {"ASN":"13335","Name":"CLOUDFLARENET"}

curl "http://localhost:10069/lookup/country?ip=8.8.8.8"
> {"Country":{"Code":"US","ID":6252001,"EuropeanUnion":false},"Continent":{"Code":"NA",...},...}

# Filter nested fields
curl "http://localhost:10069/lookup/country?ip=8.8.8.8&filter=Country.Code"
> "US"

curl "http://localhost:10069/lookup/city?ip=8.8.8.8&filter=Location"
> {"AccuracyRadius":1000,"Latitude":37.751,"Longitude":-97.822,"Timezone":"America/Chicago"}

# Health check
curl "http://localhost:10069/health"
> {"status":"ok"}
```

----

## Development

### Building

```bash
go build -o geoip-lookup ./src/cmd
```

### Testing

```bash
bash scripts/test.sh
```

----

## Thanks

O-X-L/geoip-lookup-service
