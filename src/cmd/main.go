package main

import (
	"flag"
	"fmt"

	"github.com/tracyhatemice/gogeoip-lookup/src/internal/cnf"
)

func welcome() {
	fmt.Printf("\n   ______           ________     __                __             \n")
	fmt.Println("  / ____/__  ____  /  _/ __ \\   / /   ____  ____  / /____  ______ ")
	fmt.Println(" / / __/ _ \\/ __ \\ / // /_/ /  / /   / __ \\/ __ \\/ //_/ / / / __ \\")
	fmt.Println("/ /_/ /  __/ /_/ // // ____/  / /___/ /_/ / /_/ / ,< / /_/ / /_/ /")
	fmt.Println("\\____/\\___/\\____/___/_/      /_____/\\____/\\____/_/|_|\\__,_/ .___/ ")
	fmt.Println("                                                         /_/      ")
	fmt.Printf("Version: %v\n", cnf.VERSION)
}

func main() {
	var listenAddr string
	var listenPort uint
	var dbType string

	flag.StringVar(&listenAddr, "l", "::", "Address to listen on (dual-stack)")
	flag.UintVar(&listenPort, "p", 10000, "Port to listen on")
	flag.StringVar(&dbType, "t", "ipinfo", "Database type to use (ipinfo or maxmind)")
	flag.StringVar(&cnf.DBCountry, "country", cnf.DBCountry, "Path to the country-database (optional)")
	flag.StringVar(&cnf.DBCity, "city", cnf.DBCity, "Path to the city-database (optional)")
	flag.StringVar(&cnf.DBASN, "asn", cnf.DBASN, "Path to the asn-database (optional)")
	flag.StringVar(&cnf.DBPrivacy, "privacy", cnf.DBPrivacy, "Path to the privacy-database (optional)")
	flag.BoolVar(&cnf.ReturnPlain, "plain", cnf.ReturnPlain, "If the result should be returned in plain text format")
	flag.Parse()

	if dbType == "maxmind" {
		cnf.CurrentDBType = cnf.DBTypeMaxMind
	} else {
		cnf.CurrentDBType = cnf.DBTypeIPInfo
	}

	welcome()
	httpServer(listenAddr, listenPort)
}
