package handlers

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func ShutDown(w http.ResponseWriter, r *http.Request) {
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		fmt.Println("host " + host + " port " + port)
		// Only allow from LocalHost
		if host == "::1" {
			os.Exit(0)
		}
	}
}
