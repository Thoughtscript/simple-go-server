package handlers

import (
	"fmt"
	"net/http"
)

func HeartBeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Heartbeat - Status: OK")
}
