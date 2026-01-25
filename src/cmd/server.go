package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/tracyhatemice/gogeoip-lookup/src/internal/cnf"
	"github.com/tracyhatemice/gogeoip-lookup/src/internal/lookup"
	"github.com/tracyhatemice/gogeoip-lookup/src/internal/util"
)

// writeError writes an error response with the given status code and message.
func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// writeResult writes the result in the configured format (JSON or plain text).
func writeResult(w http.ResponseWriter, data any) {
	if cnf.ReturnPlain {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%+v\n", data)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// getClientIP extracts the client IP from the request, checking proxy headers.
func getClientIP(r *http.Request) (string, error) {
	// Check X-Forwarded-For header (rightmost IP is the client)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if ip := net.ParseIP(strings.TrimSpace(ips[len(ips)-1])); ip != nil {
			return ip.String(), nil
		}
	}

	// Check X-Real-IP header
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		if ip := net.ParseIP(strings.TrimSpace(realIP)); ip != nil {
			return ip.String(), nil
		}
	}

	// Fall back to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if ip := net.ParseIP(host); ip != nil {
		if ip.String() == "::1" {
			return "127.0.0.1", nil
		}
		return ip.String(), nil
	}

	return "", errors.New("IP not found")
}

// handleLookup handles GET /lookup/{type} requests.
func handleLookup(w http.ResponseWriter, r *http.Request) {
	lookupType := r.PathValue("type")
	ipStr := r.URL.Query().Get("ip")
	filterStr := r.URL.Query().Get("filter")

	// Use client IP if not provided
	if ipStr == "" {
		clientIP, err := getClientIP(r)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Could not determine client IP")
			return
		}
		ipStr = clientIP
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		writeError(w, http.StatusBadRequest, "Invalid IP address")
		return
	}

	lookupFn := lookup.Funcs()[lookupType]
	if lookupFn == nil {
		writeError(w, http.StatusBadRequest, "Invalid lookup type")
		return
	}

	data, err := lookupFn(ip)
	if err != nil {
		log.Printf("Lookup error: type=%s ip=%s error=%v", lookupType, ipStr, err)
		writeError(w, http.StatusInternalServerError, "Lookup failed")
		return
	}

	// Apply filter if specified
	if filterStr != "" {
		filteredData := data
		for _, key := range strings.Split(filterStr, ".") {
			filteredData = util.GetMapValue(filteredData, key)
			if filteredData == nil {
				writeError(w, http.StatusBadRequest, "Invalid filter path")
				return
			}
		}
		writeResult(w, filteredData)
		return
	}

	writeResult(w, data)
}

// handleHealth handles GET /health for health checks.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func httpServer(listenAddr string, listenPort uint) {
	mux := http.NewServeMux()

	// Go 1.22+ pattern routing
	mux.HandleFunc("GET /lookup/{type}", handleLookup)
	mux.HandleFunc("GET /health", handleHealth)

	listenStr := fmt.Sprintf("%s:%d", listenAddr, listenPort)

	server := &http.Server{
		Addr:         listenStr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Listening on http://" + listenStr)
	log.Fatal(server.ListenAndServe())
}
